# logging

## Dependencies

```bash
$ go get github.com/Sirupsen/logrus
```
[Godoc Link](https://godoc.org/github.com/Sirupsen/logrus)

## Getting Started

```go
package main

import "github.com/Sirupsen/logrus"

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
}

func main() {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
```