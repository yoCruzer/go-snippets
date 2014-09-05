# declarations

Type goes after identifier:

```go
var foo int // declaration without initialization
var foo int = 42 // declaration with initialization
var foo, bar int = 42, 1302 // declare and init multiple
var foo = 42 // type omitted, will be inferred
foo := 42 // shorthand, only valid in function bodies, type is always implicit
const foo = "This is a constant"
```