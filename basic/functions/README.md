# functions

```go
// a simple function
func functionName() {}

// function with parameters
func functionName(param1 string, param2 int) {}

// multiple parameters of the same type
func functionName(param1, param2 int) {}

// return type declaration
func functionName() (int) { return 42 }

// return multiple values
func returnMulti() (int, string) { return 42, "foobar" }
var x, str = returnMulti()

// return multiple named results
func returnMulti2() (n int, s string) {
    n = 42
    s = "foobar"
    return
}
var x, str = returnMulti2()
```

## functions as values

```go
func main() {
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(3, 4))
}
```

## closures

```go
// Closures: Functions can access values that were in scope when defining the function
// adder returns an anonymous function with a closure containing the variable sum
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x // sum is declared outside, but still visible
        return sum
    }
}
```