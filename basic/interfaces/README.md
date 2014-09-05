# interfaces

Declare an interface:

```go
type Awesomizer interface {
    Awesomeize() string
}
```

Types do not declare to implement interfaces:

```go
type Foo struct {
	// ...
}
```

Instead, types implicitly satisfy an interface if they implement all required methods

```go
func (foo Foo) Awesomeize() string {
    return "Awesome!"
}
```