/*

A pointer is a variable that stores the memory address of another variable.

In simple terms, instead of holding the value, it holds the location where the value is stored.

--------------------------------------------------------------

& → Address-of operator → gives the memory address of a variable.

* → Dereference operator → gets or sets the value at the memory address.

--------------------------------------------------------------
WHY USE POINTER ?
1. Modify a value inside a function (pass by reference)
2. Avoid copying large structs/arrays (performance)
3. Share data efficiently

---------------------------------------------------------------

*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func increment(a int) {
	a = a + 1
}

func incrementByReference(a *int) {
	// now here a contain the address of x
	// so we done the operation on the origin x
	fmt.Println(a)
	*a = *a + 1 // dereference to change the value

}

func main() {

	x := 10    // normal variabel
	var p *int // p is a pointer to an int
	p = &x     // &x gets the memory address of the x and assign to p
	// now p contain the memory address of x

	fmt.Println(p)
	fmt.Println(*p) // print the value of the pointer

	// -----------------------------------------------------
	// without pointer
	increment(x) // here it does not change the value of the original x but it return the modified
	// but we want to change the value of x

	incrementByReference(&x)

	fmt.Println(x)

	// -----------------------------------------------------
	// POINTER TO STRUCT
	person := Person{
		Name: "Annu",
	}

	ptr := &person

	fmt.Println("person address: ", ptr)
	fmt.Println(ptr.Name) // You can use ptr.Field directly → "Alice"

	ptr.Age = 30
	fmt.Println(person.Age) // 30 → struct is modified

	// -----------------------------------------------------
	// POINTER TO POINTER
	y := 100

	p1 := &y

	p2 := &p1

	fmt.Println(**p2) // 100 → dereference twice

	// -----------------------------------------------------
	// NIL POINTER
	var p3 *int

	fmt.Println(p3)

	if p3 != nil {
		fmt.Println(*p3)
	}

	// =======================================================
	// 					COMMON POINTER PATTERN
	// =======================================================

}
