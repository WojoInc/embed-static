# embed-static

Generator wrapping `github.com/gin-contrib/static` to provide
the ability to serve static filesystems from an embedded filesystem.

## Background

Inspired by [observIQ/embeddable-react-final](https://github.com/observIQ/embeddable-react-final) and its associated
article. Goal was to take this idea, and package it into a reusable format via the `//go:generate` directive.

## Disclaimer

This project is still very new, and should be considered as a PoC/learning opportunity.

## Usage

Given a project structure like so:
```
mywebapp
  |- main.go
  |- go.mod
  |- ui
     |- build
         |- index.html
         |- js
            |- ...
     |- src
        |- App.js
```

You could then create a `ui.go` file under `ui/` by adding a
`//go:generate` directive to your main.go.

This follows the syntax:

`//go:generate github.com/wojoinc/embed-static/cmd/fs <package name> <dest> <dir to embed> <file to use as index>`

For example, in `main.go`:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/wojoinc/mywebapp/ui"
)

//go:generate ../embed-static/embed-static ui ui/ui.go build ui/build/index.html

func main() {
    r := gin.Default()
    ui.AddEmbeddedRoutes(r, "/")
    r.Run()
}

```

## TODO

- Fill out rest of the README
- Add examples
- Add tests