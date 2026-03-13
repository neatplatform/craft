package askit

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/bgentry/speakeasy"
	"github.com/mattn/go-isatty"
)

// Asker is the interface for prompting users and reading inputs.
type Asker interface {
	Output(string)
	Ask(string) (string, error)
	AskSecret(string) (string, error)
}

// asker implements the Asker interface.
type asker struct {
	sync.Mutex

	reader io.ReadCloser
	writer io.WriteCloser
}

// New creates a new Asker instance.
// This is a concurrent-safe Asker and can be used across multiple Go routines.
func NewAsker() Asker {
	return &asker{
		reader: os.Stdin,
		writer: os.Stdout,
	}
}

func (a *asker) Output(message string) {
	a.Lock()
	defer a.Unlock()

	_, _ = fmt.Fprintln(a.writer, message)
}

func (a *asker) Ask(prompt string) (string, error) {
	return a.ask(prompt, false)
}

func (a *asker) AskSecret(prompt string) (string, error) {
	return a.ask(prompt, true)
}

func (a *asker) ask(prompt string, secret bool) (string, error) {
	a.Lock()
	defer a.Unlock()

	if _, err := fmt.Fprint(a.writer, prompt+" "); err != nil {
		return "", err
	}

	// Listen for OS interrupt signals (e.g. Ctrl+C) so we can abort and return an error.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	errCh := make(chan error, 1)
	lineCh := make(chan string, 1)

	// Read user input in a goroutine to avoid blocking the singal handling.
	go func() {
		var line string
		var err error

		if secret && isatty.IsTerminal(os.Stdin.Fd()) {
			line, err = speakeasy.Ask("")
		} else {
			r := bufio.NewReader(a.reader)
			line, err = r.ReadString('\n')
		}

		if err != nil {
			errCh <- err
			return
		}

		lineCh <- strings.TrimRight(line, "\r\n")
	}()

	select {
	case err := <-errCh:
		return "", err
	case line := <-lineCh:
		return line, nil
	case <-sigCh:
		_, _ = fmt.Fprintln(a.writer)
		return "", errors.New("interrupted")
	}
}
