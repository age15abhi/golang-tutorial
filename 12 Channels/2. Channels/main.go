/*

🔹 1. What is a Channel?

A channel is like a pipe that connects goroutines so they can safely send and receive data.
Think of it as a communication line:


- One goroutine sends data into the channel.

- Another goroutine receives data from the channel.

- SChannels are type-safe → a chan int can only carry int.

---------------------------------------------------------------

🔹 2. Sending and Receiving

- Send → ch <- value
- Receive → x := <-ch

✅ Key rule: Send and receive block until the other side is ready.

- If you send into a channel but no one is receiving → program waits (blocks).

- If you receive but no one is sending → program waits too.

This blocking nature makes channels great for synchronization.


---------------------------------------------------------------

🔹 3. Buffered Channels

By default, channels are unbuffered → sender waits until receiver is ready.

- Sending only blocks when the buffer is full.
- Receiving only blocks when the buffer is empty.

Note: ⚡ Use buffered channels when you want asynchronous communication.

---------------------------------------------------------------

🔹 4. Closing Channels

a. Sometimes, you want to signal that no more data will be sent.

syntax -
close(ch)

b. Receivers can detect closure using:
syntax - v, ok := <-ch

- If ok == true → value received.
- If ok == false → channel is closed and empty.


---------------------------------------------------------------

🔹 5. select Statement (Channel Multiplexing)

select lets you wait on multiple channels at once.
This is super powerful for concurrent tasks.


select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
case msg := <-ch2:
    fmt.Println("Got from ch2:", msg)
default:
    fmt.Println("No message")
}

Note: Real use: handling multiple requests, timeouts, cancellations.

---------------------------------------------------------------

🔹 7. Directional Channels

You can restrict channel use:
- chan<- T → send-only
- <-chan T → receive-only

func send(ch chan<- int) {
    ch <- 42
}

func receive(ch <-chan int) {
    fmt.Println(<-ch)
}


---------------------------------------------------------------

🔹 8. Deadlocks

⚠️ Deadlock happens when all goroutines are waiting but no one can proceed.

ch := make(chan int)
ch <- 1 // deadlock, no receiver

Go runtime will panic: fatal error: all goroutines are asleep - deadlock!

---------------------------------------------------------------

🔹 9. Real-world Uses of Channels

- Handling web requests concurrently.
- Implementing timeouts and cancellation with select.
- Building worker pools for batch jobs.
- Streaming data pipelines.
- Synchronizing goroutines (instead of mutexes).



*/

package main

import "fmt"

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	// create a go routine that pull the value from in and push to out
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	// run the for loop to enter the values in the channel
	// this is goroutine thread
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func main() {
	ch := make(chan int) // channel carring int

	go func() {
		ch <- 12345
	}()

	channelNumber := <-ch
	fmt.Println(channelNumber)

	// --------------------------------------------------------
	// But you can add a buffer (like a queue):
	ch1 := make(chan int, 2) // capacity = 2

	ch1 <- 1
	ch1 <- 2
	close(ch1) // closing channel

	for v := range ch1 {
		fmt.Println(v)
	}

	// --------------------------------------------------------
	jobs := make(chan int, 5)    // capacity 5 // send
	results := make(chan int, 5) // capacity 5 // receive

	for w := 1; w <= 3; w++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for j := range jobs {
				results <- j * 2 // 2 ,4 ,6 , 8 ,10
			}
		}(w, jobs, results)

		// result -> 2 ,4 ,6 , 8 ,10
	}

	for j := 1; j <= 5; j++ {
		jobs <- j // 1 ,2 ,3 , 4,5
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}

	// --------------------------------------------------------
	// Pipeline (data flows step by step):

	ch2 := gen(2, 3, 4) // this function take value and return channel ; 2,3,4 -> out -> sq -> out -> print

	// now enter the chanel value in the sq function
	out := sq(ch2)

	for o := range out {
		fmt.Println(o)
	}
}
