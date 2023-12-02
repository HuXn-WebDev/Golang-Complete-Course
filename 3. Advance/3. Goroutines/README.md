# Goroutines

### Goroutine is a lightweight, independently running function that allows concurrent execution in a program. It's like a mini-thread managed by the Go runtime, making it easy to run multiple tasks simultaneously. Goroutines are a key feature for concurrent programming in the Go language.

# Characteristics

### 1. Concurrency: Goroutines let you do multiple things at the same time in a straightforward way. It's like juggling several tasks without dropping any balls.

### 2. Low Overhead: Goroutines are like very lightweight helpers that don't take up much space. You can have lots of them in your program without it getting too heavy.

### 3. Concurrency with Share Memory: Goroutines communicate and synchronize using shared memory (via channels), making it easy to coordinate the execution of concurrent tasks.

### 4. Language-Level Support: Goroutines are like little workers provided by the Go language itself. You don't have to worry about complicated details because the Go language helps you manage them easily. It's like having a team of workers that already know how to cooperate.

```
go funcName()

```
