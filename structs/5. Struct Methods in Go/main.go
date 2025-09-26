/*

🟢 What is a Struct Method in Go?

A method in Go is just a function with a receiver.

A receiver is like a parameter that ties the function to a specific struct type.

This allows the function to operate on the struct’s data (fields).

Think of it as functions that belong to structs — similar to methods in OOP, but Go does it with receivers instead of classes.

*/

package main

import "fmt"

// 📌 Example 1: Function vs Method

type User struct {
	name string
}

// first create the normal function
func greet(u User) {
	fmt.Println("Hello ", u.name)
}

// method (receiver is User)
func (u User) greet() {
	fmt.Println("Hello ", u.name)
}

// 📌 Example 2: Value Receiver vs Pointer Receiver

func main() {
	// create the struct
	user := User{
		name: "abhishek",
	}

	greet(user)  // calling normal function
	user.greet() // calling method of struct
}
