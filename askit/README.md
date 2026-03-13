# askit

This package lets you declaratively prompt users for input using struct tags.
Annotate your Go struct fields with `ask` tags, then hand the struct to an `Asker`
which reads values from *stdin* and get back a fully populated struct.

## Quick Start

```go
package main

import (
  "fmt"
  "os"

  "github.com/neatplatform/craft/askit"
)

func main() {
  asker := askit.NewAsker()

  info := struct {
    Name  string `ask:"any, your full name"`
    Email string `ask:"email, your email address"`
    Token string `ask:"secret, your access token"`
  }{
    Name: "Jane Doe",
  }

  err := askit.Ask(&info, asker)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", info)
}
```

## Supported Types

  - `string`, `*string`, `[]string`
  - `bool`, `*bool`, `[]bool`
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `*int`, `*int8`, `*int16`, `*int32`, `*int64`
  - `[]int`, `[]int8`, `[]int16`, `[]int32`, `[]int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`
  - `*uint`, `*uint8`, `*uint16`, `*uint32`, `*uint64`
  - `[]uint`, `[]uint8`, `[]uint16`, `[]uint32`, `[]uint64`
  - `float32`, `float64`
  - `*float32`, `*float64`
  - `[]float32`, `[]float64`
  - `byte`, `*byte`, `[]byte`
  - `rune`, `*rune`, `[]rune`
  - `url.URL`, `*url.URL`, `[]url.URL`
  - `time.Duration`, `*time.Duration`, `[]time.Duration`
  - `regexp.Regexp`, `*regexp.Regexp`, `[]regexp.Regexp`

Nested structs are supported at any depth.

Regular expressions use [POSIX Basic Regular Expression](https://en.wikibooks.org/wiki/Regular_Expressions/POSIX_Basic_Regular_Expressions) syntax.
