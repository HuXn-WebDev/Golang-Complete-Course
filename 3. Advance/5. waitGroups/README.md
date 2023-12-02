# Wait Groups

### Imagine you have a group of friends working on different tasks, and you want to make sure everyone finishes before moving on. A wait group is like a coordinator or a counter that helps you wait for all your friends to complete their tasks.

#### Add Friends to the Group:

```go
// You add your friends (goroutines) to the wait group before they start their tasks.
var wg sync.WaitGroup
```

#### Friends Start Tasks:

```go
// Your friends start working on their tasks (goroutines).
wg.Add(1)  // Adding one friend to the group
go func() {
    // Friend's task
    defer wg.Done()  // When the task is done, tell the wait group
}()
```

#### Wait for Friends to Finish:

```go
// You wait for your friends to finish their tasks.

wg.Wait()  // Wait until all friends are done
```

#### The wait group is like a helpful friend that keeps track of tasks. When a friend finishes their task, they say, "I'm done," and the wait group listens. You wait until all your friends say they're done before moving on. It helps ensure that everything is completed before you proceed with the next steps in your program.
