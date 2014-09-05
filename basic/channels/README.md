# channels

```go
ch := make(chan int) // create a channel of type int
ch <- 42             // send a value to the channel ch
v := <-ch            // receive a value from ch

// non-buffered channels block. read blocks when no value is available
// create a buffered channel. writing to buffered channels does not block
ch := make(chan int, 100)
close(c) // close the channel (only sender should close)

// read from channel and test if it has been closed
v, ok := <-ch

// if ok is false, channel has been closed
// read from channel until it is closed
for i := range ch {
    fmt.Println(i)
}

// select blocks on multiple channel operations, if one unblocks, the corresponding case is executed
func doStuff(channelOut, channelIn, chan int) {
    select {
    case channelOut <- 42:
        fmt.Println("We could write to channelOut")
    case x := <- channelIn:
        fmt.Println("We could read from channelIn")
    }
}
```