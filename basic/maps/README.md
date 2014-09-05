# maps

```go
var m map[string]int
m = make(map[string]int)

m["key"] = 42
fmt.Println(m["key"])

delete(m, "key")

elem, ok = m["key"] // test if key is present and retrieve it if so

// map literal
var m = map[string]Vertex{
    "Bell Labs": {40.6, -74.3},
    "Google": {37.4, -122.0},
}
```