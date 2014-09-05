# websocket

## Dependencies

```bash
$ go get code.google.com/p/go.net/websocket
```
[Godoc Link](http://code.google.com/p/go.net/websocket)

## Getting Started

```go
package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "io"
    "log"
)

func echoHandler(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

func main() {
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```