# packages

* Package declaration at top of every source file
* Executables are in package main
* Convention: package name == last name of import path (eg import path math/rand => package rand)
* Upper case identifier: exported (visible from other packages)
* Lower case identifier: private (not visible from other packages)

```go
package main

// ...
```