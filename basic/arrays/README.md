# arrays

```go
var a [10]int // declare an int array with length 10

a[3] = 42 // set elements

i := a[3] // read elements

// declare and init
a := [2]int{1, 2}
a := [...]int{1, 2} // elipsis -> compiler figures out length
```

# slices

```go
var a []int // declare a slice, length is unspecified

var a = []int {1, 2, 3, 4} // declare and init a slice

a := []int{1, 2, 3, 4} // shorthand

var b = a[lo:hi] // create a slice from index lo to hi-1

var b = a[1:4] // slice from index 1 to 3

// create a slice with make
a = make([]byte, 5, 5) // first arg length, second capacity
a = make([]byte, 5) // capacity is optional
```

## operations on arrays and slices

```go
len(a) gives the length of the array / slice.

// loop over an array / slice
for i, e := range a {
    // i is the index, e is the element
}

// if you only need e
for _, e := range a {
    // e is the element
}

// and if you only need the index
for i := range a {
}
```