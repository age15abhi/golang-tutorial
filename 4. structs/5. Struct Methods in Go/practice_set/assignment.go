/* ====================================================

📝 Assignment for You

👉 Build this step by step:

Task

1. Create a struct Rectangle with fields:

	- width (float64)
	- height (float64)

2. Add two methods:

	- Area() → returns the area of the rectangle.
	- Perimeter() → returns the perimeter of the rectangle.

3. In main():

	- Create a rectangle with width = 10, height = 5.
	- Print its area and perimeter.

4. 🔥 Bonus: Write a method Scale(factor float64) (with
	pointer receiver) that multiplies the rectangle’s width and height by that factor.

	- Call it with rect.Scale(2) and show the updated 	area/perimeter.

===================================================== */

package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	rectangle := Rectangle{
		width:  10,
		height: 5,
	}

	area := rectangle.Area()
	perimeter := rectangle.Perimeter()

	fmt.Println(area)
	fmt.Println(perimeter)

}
