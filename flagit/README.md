# flagit

This package allows you to use the `flag` struct tag on your Go struct fields.
You can then read the values for those fields from the command-line arguments.

This package maps command-line flags to Go struct fields using the `flag` struct tag.
Define your configuration as a struct, call `Parse`, and the library handles the rest,
including parsing, nested structs, pointer types, and slices.

## Quick Start

```go
package main

import (
  "fmt"
  "net/url"
  "time"

  "github.com/neatplatform/craft/flagit"
)

// Spec is a struct for mapping its fields to command-line flags.
type Spec struct {
  // Flag fields
  Verbose bool `flag:"verbose"`

  // Nested fields
  Options struct {
    Port     uint16 `flag:"port"`
    LogLevel string `flag:"log-level"`
  }

  // Nested fields with prefix
  Config struct {
    Timeout   time.Duration `flag:"timeout"`
    Endpoints []url.URL     `flag:"endpoints"`
  } `flag:"config-"`
}

func main() {
  spec := new(Spec)
  if err := flagit.Parse(spec, false); err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", spec)
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
Assigning a `flag` tag to a struct field uses its value as a prefix for all descendant flags.

Regular expressions use [POSIX Basic Regular Expression](https://en.wikibooks.org/wiki/Regular_Expressions/POSIX_Basic_Regular_Expressions) syntax.
