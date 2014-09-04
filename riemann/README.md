### riemann

## Dependencies

```bash
$ go get github.com/bigdatadev/goryman
```
[Godoc Link](http://godoc.org/github.com/bigdatadev/goryman)

## Getting Started

```go
package main

import "github.com/bigdatadev/goryman"

func main() {
	c := goryman.NewGorymanClient("localhost:5555")
	err := c.Connect()
	if err != nil {
	    panic(err)
	}
	defer c.Close()

	err = c.SendEvent(&goryman.Event{
	    Service: "moargore",
	    Metric:  100,
	    Tags: []string{"nonblocking"},
	})
	if err != nil {
	    panic(err)
	}
}
```