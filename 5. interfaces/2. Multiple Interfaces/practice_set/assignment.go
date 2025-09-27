/*

📝 Assignment: Multiple Interfaces
Task:

1. Define two interfaces:

	- Flyer with method Fly()
	- Swimmer with method Swim()

2. Create a struct Duck that implements both Flyer and Swimmer.

	- Fly() should print "Duck is flying"
	- Swim() should print "Duck is swimming"

3. Create a function DoFly(f Flyer) and DoSwim(s Swimmer) that accept those interfaces.

4. In main(), create a Duck and call both DoFly and DoSwim.

*/

package main

import "fmt"

type Flyer interface {
	Fly()
}

type Swimmer interface {
	Swim()
}

type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Duck is flying")
}

func (d Duck) Swim() {
	fmt.Println("Duck is swimming")
}

func DoFly(f Flyer) {
	f.Fly()
}

func DoSwim(s Swimmer) {
	s.Swim()
}

func main() {
	d := Duck{}

	DoFly(d)
	DoSwim(d)
}
