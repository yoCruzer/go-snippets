# structs

There are no classes, only structs with methods.

```go
// declaration
type Vertex struct {
    X, Y int
}

// creating
var v = Vertex{1, 2}

// accessing members
v.X = 4

// declare methods on structs
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// call method
v.Abs()

// for mutating methods you need to use a pointer to the struct as the type
func (v *Vertex) add(n float64) {
    v.X += n
    v.Y += n
}
```