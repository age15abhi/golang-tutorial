package main

import "fmt"

func main() {

	//	Task 1: Classic for loop
	// Write a for loop that prints numbers from 1 to 10.
	// Inside the loop, skip the number 5 using continue.
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//	Task 2: While-style for loop
	// Use a while-style for loop to print numbers from 10 down to 1.
	// Stop the loop when it reaches 1 using break.
	num := 10

	for num >= 1 {
		fmt.Println(num)
		num--
	}

	// Task 3: Infinite loop
	// Create an infinite loop that prints "Running...".
	// Stop the loop after printing 5 times using break.

	count := 1
	for {
		fmt.Println("count:", count)
		if count == 5 {
			break
		}
		count++

	}

	//	Task 4: Loop with range over a slice
	//
	// Create a slice of integers: [2, 4, 6, 8, 10].
	// Use a for ... range loop to print:
	// Index of each element
	// Value of each element
	integers := []int{2, 4, 6, 8, 10}

	for index, value := range integers {
		fmt.Println("index:", index, "value:", value)
	}

	// 	Task 5: Loop with range over a map
	// Create a map: map[string]int{"Alice": 23, "Bob": 30, "Charlie": 27}.
	// Loop over it using range and print:
	// Key and value
	myMap := map[string]int{"Alice": 23, "Bob": 30, "Charlie": 27}

	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)
	}

	//	Task 6: Nested loop
	//
	// Use nested loops to print a multiplication table from 1 to 5.
	// Example output:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}
}
