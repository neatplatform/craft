# ui

Package `ui` provides a UI interface and default implementation for
printing styled, leveled output in command-line applications.

It supports configurable **verbosity levels** and a range of color and formatting **styles**. 

## Quick Start

```go
package main

import "github.com/neatplatform/craft/ui"

func main() {
  u := ui.New(ui.Info)
  u.Infof(ui.Green, "Hello, %s!", "World")
}
```
