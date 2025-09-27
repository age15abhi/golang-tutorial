package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {

	var s Shape
	c := Circle{
		radius: 10,
	}

	s = c

	// “I know this interface actually holds a Circle. Let me get the Circle out so I can access Circle-specific things.”

	// ok → safe check:
	// If s really contains a Circle, ok is true and c is the concrete Circle.
	// If not, ok is false and no panic occurs.
	cir, ok := s.(Circle)

	if ok {
		fmt.Println("Circle radius:", cir.radius)
	}
}

/*

Access fields or methods not in the interface

Shape only has Area().

Circle has Radius.

If you want Radius, you must get the concrete type using type assertion.

3️⃣ When type assertion is useful

Type assertion is only needed when you want something specific about the concrete type that the interface does NOT provide.

*/
