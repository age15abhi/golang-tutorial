/*
A variadic function is a function that can take zero, one, or many arguments of the same type.

👉 In Go, we declare it using ... before the type in the parameter list.

```

	func functionName(param ...Type) {
	    // param is a slice of Type
	}

```
*/
package main

import "fmt"

// variadic function
/*

👉 How it works?

numbers ...int → inside the function, numbers is treated like a slice of ints ([]int).

So in the first call, sum(1, 2, 3) → numbers = []int{1,2,3}.

*/

func sum(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println(prefix, name)
	}
}

func main() {

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum())

	// --------------------------------
	// mixed with normal params
	greet("Hello", "abhi", "annu", "ravi")

}
