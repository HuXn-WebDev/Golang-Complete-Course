# Channels

### Channels in Go are like communication pipes between different parts of your program. They let different pieces of your code talk to each other by sending and receiving messages. It's like passing notes or messages between friends to coordinate what needs to be done.

### Imagine you have two friends working on a project. One friend can put notes in the channel, and the other friend can pick up those notes. This helps them share information and work together without getting in each other's way. That's what channels do in Go—they help different parts of your program communicate smoothly.

#### Creating a channel is like making a special mailbox for messages. You can use the make function to create this mailbox. It's like setting up a mailbox with a specific type of messages it can hold.

```go
// Creating a channel for sending and receiving messages of type int
myChannel := make(chan int)

```

```go
// Sending the number 20 into the channel
myChannel <- 20
```

```go
// Receiving a message from the channel and storing it in the variable 'result'
result := <-myChannel
```

# Unbuffered Channels

### 1. An unbuffered channel has a capacity of 0.

### 2. Each send operation on an unbuffered channel blocks until there is a corresponding receive operation, and vice versa.

### 3. This creates a direct, synchronous communication between the sender and the receiver.

# Buffered Channels

### 1. A buffered channel has a specified capacity greater than 0.

### 2. It allows multiple values to be sent into the channel without an immediate corresponding receiver.

### 3. Send operations on a buffered channel block only when the buffer is full, and receive operations block only when the buffer is empty.
