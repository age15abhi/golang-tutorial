/*

You already know about sync.Mutex (a normal lock: only one goroutine can access the resource at a time).

Now, sync.RWMutex is like a more advanced lock that allows multiple readers but only one writer.

📝 Analogy (Easy Way)

Imagine you have a library:

📖 Readers (RLock): Many people can sit and read the same book at once. No problem, because reading does not change the book.

✍️ Writer (Lock): But if one person wants to write or update the book, no one else can read or write until that person finishes.

So:

RLock() → multiple goroutines can read safely.

Lock() → only one goroutine can write, and it blocks all readers and writers.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var data = 0
var rwmu sync.RWMutex

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	rwmu.RLock() // take read lock
	fmt.Printf("Reader %d: sees data = %d\n", id, data)
	time.Sleep(100 * time.Millisecond)
	rwmu.RUnlock() // ✅ release read lock
}

func writer(id int, wg *sync.WaitGroup) {
	// first of all write the defer
	defer wg.Done()

	// put the lock before writing
	rwmu.Lock() // take write lock
	fmt.Printf("Writer %d: updating data...\n", id)
	data += 1
	time.Sleep(200 * time.Millisecond)

	fmt.Printf("Writer %d: updated data = %d\n", id, data)
	rwmu.Unlock() // release write lock
}

func main() {

	// create the waiting group to keep record the wait group
	var wg sync.WaitGroup

	// start readers
	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go reader(i, &wg)
	}

	// start writer
	for i := 1; i <= 2; i++ {
		wg.Add(1)

		go writer(i, &wg)
	}
	wg.Wait()
	fmt.Println("Final data:", data)

}
