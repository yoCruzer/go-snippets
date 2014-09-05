# goroutines

Goroutines are lightweight threads managed by Go. go f(a, b) starts a new goroutine which runs f (given f is a function)

```go
func doStuff(s string) {
	// ...
}

func main() {
    go doStuff("foobar")
    go func (x int) {
        // function body goes here
    }(42)
}
```