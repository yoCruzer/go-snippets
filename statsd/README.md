# statsd

## Dependencies

```bash
$ go get github.com/quipo/statsd
```
[Godoc Link](http://godoc.org/github.com/quipo/statsd)

## Getting Started

```go
package main

import "github.com/quipo/statsd"

func main() {
    prefix := "myproject."
    statsdclient := statsd.NewStatsdClient("localhost:8125", prefix)
    statsdclient.CreateSocket()
    stats := statsd.NewStatsdBuffer(interval, statsdclient)
    defer stats.Close()

    // not buffered: send immediately
    statsdclient.Incr("mymetric", 4)

    // buffered: aggregate in memory before flushing
    stats.Incr("mymetric", 1)
    stats.Incr("mymetric", 3)
    stats.Incr("mymetric", 1)
    stats.Incr("mymetric", 1)
}
```