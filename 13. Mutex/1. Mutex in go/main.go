package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	mu.Lock() // aquire lock
	counter++ // critical section
	fmt.Println(counter)
	mu.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		// this create the go routine
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter", counter)
}
