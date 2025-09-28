/*

In Go, the only looping construct is the for loop.
There is no while or do-while like in other languages — but for can be used in all those ways.

*/

package main

import "fmt"

func main() {

	// 1️⃣ Classic for loop (like C/Java)
	for i := 0; i < 5; i++ {
		fmt.Println("iteration:", i)
	}

	//  2️⃣ While-style for loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 3️⃣ Infinite loop
	// for {
	// 	fmt.Println("This will run forever!")
	// }
	count := 0
	for {
		count++
		fmt.Println("Count: ", count)
		// check for the count == 5

		if count >= 5 {
			break
		}
	}

	// 4 Loop over arrays, slices, maps, strings (range)
	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}

	// if you do not need index
	for _, value := range nums {
		fmt.Println("Only value: ", value)
	}

	// 5. Iterate on the map
	myMap := map[string]int{"a": 1, "b": 2}
	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)
	}
}
