/**

📝 Assignment
Task:

1. Create an interface Animal with method Speak() string.

2. Create two structs:
	- Dog → Speak() returns "Woof"
	- Cat → Speak() returns "Meow"

3. Write a function Identify(a Animal) that:

	- Calls a.Speak()
	- Uses type assertion to check if a is a Dog or a Cat, and prints:
		- "This is a Dog"
		- "This is a Cat"

*/

package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

func Identify(a Animal) {
	fmt.Println("Animal says:", a.Speak()) // print Speak()
	switch a.(type) {
	case Cat:
		fmt.Println("this is Cat")
	case Dog:
		fmt.Println("This is a Dog")
	default:
		fmt.Println("Unknown Animal")
	}

}

func main() {

	d := Dog{}
	c := Cat{}

	Identify(d)
	Identify(c)
}
