package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFg256Color(t *testing.T) {
	tests := []struct {
		name          string
		code          int
		expectedStyle Style
	}{
		{
			name:          "Default",
			code:          999,
			expectedStyle: Style{38, 5, 0},
		},
		{
			name:          "Orange",
			code:          214,
			expectedStyle: Style{38, 5, 214},
		},
		{
			name:          "RedOrange",
			code:          202,
			expectedStyle: Style{38, 5, 202},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStyle, Fg256Color(tc.code))
		})
	}
}

func TestBg256Color(t *testing.T) {
	tests := []struct {
		name          string
		code          int
		expectedStyle Style
	}{
		{
			name:          "Default",
			code:          999,
			expectedStyle: Style{48, 5, 7},
		},
		{
			name:          "BrightGreen",
			code:          82,
			expectedStyle: Style{48, 5, 82},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStyle, Bg256Color(tc.code))
		})
	}
}

func TestFgTrueColor(t *testing.T) {
	tests := []struct {
		name          string
		rgb           int
		expectedStyle Style
	}{
		{
			name:          "Default",
			rgb:           0x000000,
			expectedStyle: Style{38, 2, 0, 0, 0},
		},
		{
			name:          "OK",
			rgb:           0x8ECAE6,
			expectedStyle: Style{38, 2, 142, 202, 230},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStyle, FgTrueColor(tc.rgb))
		})
	}
}

func TestBgTrueColor(t *testing.T) {
	tests := []struct {
		name          string
		rgb           int
		expectedStyle Style
	}{
		{
			name:          "Default",
			rgb:           0xFFFFFF,
			expectedStyle: Style{48, 2, 255, 255, 255},
		},
		{
			name:          "OK",
			rgb:           0x023047,
			expectedStyle: Style{48, 2, 2, 48, 71},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStyle, BgTrueColor(tc.rgb))
		})
	}
}

func TestStyle_sprintf(t *testing.T) {
	tests := []struct {
		name           string
		s              Style
		format         string
		args           []interface{}
		expectedString string
	}{
		{
			name:           "Bold",
			s:              Style{Bold},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedString: "\x1b[1mHello, World!\x1b[0m",
		},
		{
			name:           "FgGreen",
			s:              Style{FgGreen},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedString: "\x1b[32mHello, World!\x1b[0m",
		},
		{
			name:           "BgBlue",
			s:              Style{BgBlue},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedString: "\x1b[44mHello, World!\x1b[0m",
		},
		{
			name:           "MixStyle",
			s:              Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedString: "\x1b[43;35;1;5mHello, World!\x1b[0m",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s.sprintf(tc.format, tc.args...)

			assert.Equal(t, tc.expectedString, s)
		})
	}
}
