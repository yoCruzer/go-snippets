# http client

## Getting Started

```go
package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person := &Person{"Jonh", 27}

	buf, _ := json.Marshal(person)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8080", "application/json", body)

	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}
```