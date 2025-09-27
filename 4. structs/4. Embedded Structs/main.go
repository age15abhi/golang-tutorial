/*

🟢 What are Embedded Structs?

In Go, embedded structs mean putting one struct inside another struct without a field name.
This is sometimes called composition (Go doesn’t have inheritance like Java or C++).

➡️ By embedding, the fields and methods of the inner struct become directly accessible from the outer struct.

*/

package main

import "fmt"

// Example 1 : Embedded Struct

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	Address // embedded struct (no field name)
}

// 📌 Example 2: Methods with Embedded Structs

type User struct {
	name string
}

func (u User) greet() {
	fmt.Println("Hello, ", u.name)
}

type Admin struct {
	User
	level int
}

func main() {

	p := Person{
		name: "abhishek",
		age:  21,
		Address: Address{
			city:  "jaipur",
			state: "Rajathan",
		},
	}

	a := Admin{
		User:  User{name: "abhishek"},
		level: 1,
	}

	// in the greet function we pass the this of the User
	a.greet()

	// 👇 You can access Address fields directly
	fmt.Println(p.city)
}
