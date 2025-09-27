/*

1. What is an Interface in Go?

In Go, an interface is a type that defines a set of method signatures (behavior) but does not provide implementation.

👉 You can think of it as:

“If a type implements all the methods listed in an interface, it implicitly satisfies that interface.”

Unlike Java or C#, in Go you don’t need to explicitly declare implements.

🔑 The Purpose of an Interface (in simple words)

👉 An interface is for writing code that works with many different types, without caring about their details.

It’s about flexibility and reusability.

------------------------------------------------------

📍 Real-life analogy

Imagine you own a company.

	- You don’t care who is working for you (a human or a robot).
	- You only care that they can do the job (e.g., Work()).

So instead of writing:

	- HirePerson(p Person)
	- HireRobot(r Robot)

You just write:

	- Hire(worker Worker)

Now, anyone (person, robot, AI) who knows how to Work() can be hired.

*/

package main

import "fmt"

// -------------------------------------------------

type Worker interface {
	work()
}

type Person struct {
	name string
}

func (p Person) work() {
	fmt.Println(p.name)
}

// -------------------------------------------------
// this is the main function it can used by any
type Shape interface {
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// so i can define the same function again and again for the differnt struct
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func main() {
	var w Worker
	p := Person{
		name: "Abhi",
	}

	w = p

	w.work()

	// ---------------------------------

	r := Rectangle{
		Width:  10,
		Height: 5,
	}
	c := Circle{Radius: 7}

	PrintArea(r) // Rectangle is a Shape
	PrintArea(c) // Circle is a Shape
}
