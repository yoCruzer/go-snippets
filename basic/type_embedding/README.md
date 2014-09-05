# type embedding

Here are two interfaces:

```go
package bufio

type Reader interface {
    Read(p []byte)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Use type embedding to embed two interfaces to form a new one:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

Combine a reader and a writer into a struct using embedding:

```go
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```