### memcached

## Dependencies

```bash
$ go get github.com/bradfitz/gomemcache/memcache
```
[Godoc Link](http://godoc.org/github.com/bradfitz/gomemcache/memcache)

## Getting Started

```go
package main

import "github.com/bradfitz/gomemcache/memcache"

func main() {
     mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
     mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

     it, err := mc.Get("foo")
     ...
}
```