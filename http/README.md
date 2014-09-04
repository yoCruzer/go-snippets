### http

## Dependencies

```bash
$ go get github.com/go-martini/martini
```
[Godoc Link](http://godoc.org/github.com/go-martini/martini)

## Getting Started

```go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```