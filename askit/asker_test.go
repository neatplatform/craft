package askit

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nopWriteCloser struct {
	io.Writer
}

func (w *nopWriteCloser) Close() error {
	return nil
}

func TestNew(t *testing.T) {
	a := NewAsker()
	assert.NotNil(t, a)

	aa, ok := a.(*asker)
	assert.True(t, ok)
	assert.NotNil(t, aa.reader)
	assert.NotNil(t, aa.writer)
}

func TestUI_Output(t *testing.T) {
	tests := []struct {
		name           string
		message        string
		expectedOutput string
	}{
		{
			name:           "OK",
			message:        "Hello, World!",
			expectedOutput: "Hello, World!\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			a := &asker{
				writer: &nopWriteCloser{Writer: buf},
			}

			a.Output(tc.message)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_Ask(t *testing.T) {
	tests := []struct {
		name           string
		prompt         string
		input          string
		expectedOutput string
		expectedAnswer string
		expectedError  string
	}{
		{
			name:           "Success",
			prompt:         "Enter name:",
			input:          "Nietzsche\n",
			expectedOutput: "Enter name: ",
			expectedAnswer: "Nietzsche",
		},
		{
			name:          "EOF",
			prompt:        "Enter name:",
			input:         "",
			expectedError: "EOF",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			a := &asker{
				reader: io.NopCloser(strings.NewReader(tc.input)),
				writer: &nopWriteCloser{Writer: buf},
			}

			ans, err := a.Ask(tc.prompt)

			if tc.expectedError != "" {
				assert.Empty(t, ans)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAnswer, ans)
				assert.Equal(t, tc.expectedOutput, buf.String())
			}
		})
	}
}

func TestUI_AskSecret(t *testing.T) {
	tests := []struct {
		name           string
		prompt         string
		input          string
		expectedOutput string
		expectedAnswer string
		expectedError  string
	}{
		{
			name:           "Success",
			prompt:         "Enter password:",
			input:          "1234\n",
			expectedOutput: "Enter password: ",
			expectedAnswer: "1234",
		},
		{
			name:          "EOF",
			prompt:        "Enter password:",
			input:         "",
			expectedError: "EOF",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			a := &asker{
				reader: io.NopCloser(strings.NewReader(tc.input)),
				writer: &nopWriteCloser{Writer: buf},
			}

			ans, err := a.AskSecret(tc.prompt)

			if tc.expectedError != "" {
				assert.Empty(t, ans)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAnswer, ans)
				assert.Equal(t, tc.expectedOutput, buf.String())
			}
		})
	}
}
