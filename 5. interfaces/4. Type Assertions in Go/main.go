/*

🔑 What is a Type Assertion in Go?

When you use an interface in Go, the variable doesn’t just store the value, it also stores the type of the value.

	- A type assertion lets you pull out the original value from the interface.

👉 In short:

	- value := interfaceVar.(ConcreteType)

*/

package main

import "fmt"

// Example 1: Simple type assertion

func main() {

	var i interface{} = 42 // empty interface can hold any value so we not need define the type

	// Type assertion
	// s := i.(string) // get back the assertion
	// fmt.Println("String value:", s)

	// safe type assertion
	s, ok := i.(string)

	if ok {
		fmt.Println("string value:", s)
	} else {
		fmt.Println("Not a string!") // print this
	}

	n, ok := i.(int)

	if ok {
		fmt.Println("Number value:", n)
	}

	// If you try wrong type → panic
	// n := i.(int)  // ❌ panic: interface {} is string, not int

}
