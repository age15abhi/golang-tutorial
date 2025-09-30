In Go, concurrency means dealing with many things at once — not necessarily running them all at the same exact time (that’s parallelism).

Go makes concurrency a first-class feature through goroutines and channels.

1. Goroutines
- Lightweight threads managed by the Go runtime.
- Created using go keyword

Example:

```go

package main

import (
    "fmt"
    "time"
)

func task(name string) {
    for i := 1; i <= 3; i++ {
        fmt.Println(name, ":", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go task("Task A") // run concurrently
    go task("Task B")

    // Give time for goroutines to run
    time.Sleep(2 * time.Second)
    fmt.Println("Main finished")
}

```

2. Channels:
- Safe way for goroutines to communicate.
- They prevent race conditions without using locks.

```go
package main

import "fmt"

func worker(ch chan string) {
    ch <- "work done"
}

func main() {
    ch := make(chan string)
    go worker(ch)

    msg := <-ch // receive message
    fmt.Println(msg)
}
```

3. Concurrency vs Parallelism

- Concurrency = structuring your program so tasks can be executed independently.

- Parallelism = actually executing tasks at the same time (depends on CPU cores).

Go is designed for concurrency; parallel execution is possible if your machine has multiple cores.


✅ Real-world Use Cases of Concurrency in Go

- Handling multiple HTTP requests in a web server (each request in its own goroutine).

- Performing I/O-bound operations (reading files, databases, APIs) without blocking.

- Running background tasks (e.g., cleanup, logging) while main program continues.

- Building pipelines where data flows through channels step by step.