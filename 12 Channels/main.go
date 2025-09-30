package main

import (
	"fmt"
	"time"
)

func task(name string) {

	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func worker(ch chan string) {
	ch <- "work done"
}

func main() {
	go task("Task A")
	go task("Task B")

	// Give time for goroutines to run
	time.Sleep(2 * time.Second)
	fmt.Println("Main finished")

	ch := make(chan string)

	go worker(ch)

	msg := <-ch
	fmt.Println(msg)

}
