# shell

A Go package for running shell commands with minimal boilerplate,
supporting custom arguments, environment variables, and working directories.
It is designed from the ground up for **brevity** and **testability**.

## Why use this instead of `os/exec`?

The standard `os/exec` is powerful but verbose — capturing output, wiring up environment variables,
and checking exit codes all require boilerplate that obscures what your code is actually doing.

This package wraps that complexity behind a clean, minimal API so a shell command is a single readable line.
It also introduces `RunnerFunc`, a function type that lets you build pre-configured,
injectable command runners and swap them out for mocks in unit tests with no extra interfaces required.

## Quick Start

**Running a command:**

```go
package main

import (
  "context"
  "fmt"

  "github.com/neatplatform/craft/shell"
)

func main() {
  exitcode, stdout, err := shell.Run(context.Background(), "date", "-u")
  if err != nil {
    panic(err)
  }

  fmt.Printf("[%d] %s\n", exitcode, stdout)
}
```

**Running a command with environment variables:**

```go
package main

import (
  "context"
  "fmt"

  "github.com/neatplatform/craft/shell"
)

func main() {
  opts := shell.RunOptions{
    Environment: map[string]string{
      "GREETING": "Hello, World!",
    },
  }

  exitcode, stdout, err := shell.RunWith(context.Background(), opts, "printenv", "GREETING")
  if err != nil {
    panic(err)
  }

  fmt.Printf("[%d] %s\n", exitcode, stdout)
}
```

**Building injectable commands for testing:**

```go
package main

import (
  "context"
  "fmt"

  "github.com/neatplatform/craft/shell"
)

type service struct {
  funcs struct {
    ls shell.RunnerWithFunc
  }
}

func newService() *service {
  s := new(service)
  s.funcs.ls = shell.RunnerWith("ls")
  s.funcs.ls = s.funcs.ls.WithArgs("-a")
  return s
}

func (s *service) list(path string) (string, error) {
  opts := shell.RunOptions{WorkingDir: path}
  _, stdout, err := s.funcs.ls(context.Background(), opts)
  return stdout, err
}

func main() {
  s := newService()
  out, err := s.list("/opt")
  if err != nil {
    panic(err)
  }

  fmt.Println(out)
}
```

```go
package main

import (
  "context"
  "testing"

  "github.com/neatplatform/craft/shell"
)

func TestService_List(t *testing.T) {
  t.Run("Success", func(t *testing.T) {
    s := new(service)
    s.funcs.ls = func(context.Context, shell.RunOptions, ...string) (int, string, error) {
      return 0, "foo bar", nil
    }

    out, err := s.list("/test")
    if out != "foo bar" || err != nil {
      t.Fail()
    }
  })
}
```
