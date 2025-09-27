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

/* ====================================================
  📌 Example 2: Value Receiver vs Pointer Receiver
➡️ Use value receiver when method doesn’t modify struct.
➡️ Use pointer receiver when method should update struct or avoid copying large structs.
====================================================== */

// 1. Value Receiver (copy)
func (u User) changeNameOnlyCopy(newName string) {
	u.name = newName // only change the copy
}

// 2. Pointer Receiver (reference)
func (u *User) changeNameOriginal(newName string) {
	u.name = newName // modifies original
}

/* ====================================================
  📌 Example 3: Methods on Embedded Structs
====================================================== */

type Engine struct {
	horsepower int
}

// now can i call the function with car method
func (e Engine) info() {
	fmt.Println("HorsePower ==> ", e.horsepower)
}

type Car struct {
	brand string
	Engine
}

// ---------------- MAIN FUNCTION -----------------------

func main() {
	// create the struct
	user := User{
		name: "abhishek",
	}

	greet(user)  // calling normal function
	user.greet() // calling method of struct

	user.changeNameOnlyCopy("john")
	fmt.Println(user.name) // still "Abhi"

	user.changeNameOriginal("john")
	fmt.Println(user.name) // // now "John"

	car := Car{
		brand:  "Tesla",
		Engine: Engine{horsepower: 500},
	}

	// run the info function
	car.info()
}
