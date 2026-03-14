package ui_test

import (
	"time"

	"github.com/neatplatform/craft/ui"
)

func Example() {
	u := ui.New(ui.Info)

	u.Infof(ui.FgTrueColor(0xFF0054), "Now: %s", time.Now())
}
