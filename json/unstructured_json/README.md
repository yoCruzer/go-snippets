# Unstructured JSON

## Dependencies

JSON support is already a part of Go, so no additional dependencies are required.

## Getting Started

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// define the type as a generic map
type Config map[string]interface{}

func main() {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot read configuration file", filename)
		os.Exit(1)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("Cannot parse configuration file:", err)
		os.Exit(1)
	}

	// doSomethingUseful(c["address"], c["port"])
}
```
