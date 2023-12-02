# Select

### The select statement is used to handle multiple channels in a non-blocking way. It provides a way to wait on multiple communication operations simultaneously and execute the first case that is ready.

### Imagine you have several friends sending you messages, and you want to read whatever message comes first. The select statement is like a way to check all your messages and respond to the first one that arrives.

```go
// Whichever friend's message arrives first is the one you respond to.
select {
case message1 := <-channel1:
    fmt.Println("Received from friend 1:", message1)
case number := <-channel2:
    fmt.Println("Received from friend 2:", number)
}
```
