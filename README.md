# craft

Craft is a lightweight toolkit for building minimal, elegant command-line applications in Go.

## Quick Start

```go
package main

import (
  "github.com/neatplatform/craft/askit"
  "github.com/neatplatform/craft/ui"
)

func main() {
  a := askit.NewAsker()
  u := ui.New(ui.Info)

  inputs := struct {
    Name  string `ask:"any, your full name"`
    Email string `ask:"email, your email address"`
  }{}

  err := askit.Ask(&inputs, a)
  if err != nil {
    panic(err)
  }

  u.Infof(ui.FgTrueColor(0x006D77), "Hello, %s <%s>!", inputs.Name, inputs.Email)
}
```
