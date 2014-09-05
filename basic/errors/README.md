# errors

No exception handling, functions that might produce an error just declare an additional return value of type Error. Here is the interface:

```go
type error interface {
    Error() string
}
```

A function that might return an error:

```go
func doStuff() (int, error) {
	// ...
}

func main() {
    result, error := doStuff()
    if (error != nil) {
        // handle error
    } else {
        // use result
    }
}
```