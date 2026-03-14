[![Go Doc][godoc-image]][godoc-url]
[![Build Status][workflow-image]][workflow-url]
[![Test Coverage][codecov-image]][codecov-url]

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


[godoc-url]: https://pkg.go.dev/github.com/neatplatform/craft
[godoc-image]: https://pkg.go.dev/badge/github.com/neatplatform/craft
[workflow-url]: https://github.com/neatplatform/craft/actions/workflows/go.yml
[workflow-image]: https://github.com/neatplatform/craft/actions/workflows/go.yml/badge.svg
[codecov-url]: https://codecov.io/github/neatplatform/craft
[codecov-image]: https://codecov.io/github/neatplatform/craft/badge.svg
