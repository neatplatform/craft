package ui

import (
	"fmt"
	"strconv"
	"strings"
)

// ANSICode represents an ANSI escape code used for terminal text formatting.
type ANSICode int

const (
	Reset ANSICode = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const (
	FgBlack ANSICode = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	BgBlack ANSICode = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	Fg256 ANSICode = 38
	Bg256 ANSICode = 48
)

// Style allows multiple ANSI codes for terminal text formatting.
type Style []ANSICode

// ANSI Colors
var (
	Black   = Style{FgBlack}
	Red     = Style{FgRed}
	Green   = Style{FgGreen}
	Yellow  = Style{FgYellow}
	Blue    = Style{FgBlue}
	Magenta = Style{FgMagenta}
	Cyan    = Style{FgCyan}
	White   = Style{FgWhite}
)

// Fg256Color creates a style for a 256-color foreground color.
func Fg256Color(code int) Style {
	if code < 0 || 255 < code {
		// Default to black if out of range
		code = 0
	}

	return Style{Fg256, ANSICode(5), ANSICode(code)}
}

// Bg256Color creates a style for a 256-color background color.
func Bg256Color(code int) Style {
	if code < 0 || 255 < code {
		// Default to white if out of range
		code = 7
	}

	return Style{Bg256, ANSICode(5), ANSICode(code)}
}

// FgTrueColor creates a style for a 24-bit foreground color.
func FgTrueColor(rgb int) Style {
	if rgb < 0 || 0xFFFFFF < rgb {
		// Default to black if out of range
		rgb = 0x000000
	}

	r := (rgb >> 16) & 0xFF
	g := (rgb >> 8) & 0xFF
	b := rgb & 0xFF

	return Style{Fg256, ANSICode(2), ANSICode(r), ANSICode(g), ANSICode(b)}
}

// BgTrueColor creates a style for a 24-bit background color.
func BgTrueColor(rgb int) Style {
	if rgb < 0 || 0xFFFFFF < rgb {
		// Default to white if out of range
		rgb = 0x000000
	}

	r := (rgb >> 16) & 0xFF
	g := (rgb >> 8) & 0xFF
	b := rgb & 0xFF

	return Style{Bg256, ANSICode(2), ANSICode(r), ANSICode(g), ANSICode(b)}
}

func (s Style) sprintf(format string, a ...interface{}) string {
	const escape = "\x1b"

	codes := make([]string, len(s))
	for i, v := range s {
		codes[i] = strconv.Itoa(int(v))
	}
	sequence := strings.Join(codes, ";")

	ansiFormat := fmt.Sprintf("%s[%sm", escape, sequence)
	ansiReset := fmt.Sprintf("%s[%dm", escape, Reset)

	return ansiFormat + fmt.Sprintf(format, a...) + ansiReset
}
