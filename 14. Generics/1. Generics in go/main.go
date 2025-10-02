/*

Generics let you write functions and types that work with any data type, without repeating the same code for every type.

👉 Before Go 1.18 (released in 2022), if you wanted a function that works with both int and float64, you had to either:

Write two separate functions (AddInt, AddFloat), or

Use interface{} (but you’d lose type safety).

Now, with generics, you can write one function that works for any type safely.


*/

package main

import "fmt"

// Example 1 : Without Generics 😩 Problem → Code duplication.

func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}

// Example 2 : With Generics

func add[T int | float64](a, b T) T {
	return a + b
}

// --------------------- Restriction ------------------
type Number interface {
	int | float64
}

func Multiply[T Number](a, b T) T {
	return a * b
}

// -------------------------- generic struct --------------

// make this struct of the generic type
// type Stack struct {
// 	items []int
// }

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() T {
	last := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return last
}

func main() {
	fmt.Println(add(10, 20))   // works with int
	fmt.Println(add(3.5, 6.2)) // works with float64

	intStack := Stack[int]{}

	intStack.push(10)

	intStack.push(20)

	intStack.pop()

	fmt.Print(intStack)

	stringStack := Stack[string]{}

	stringStack.push("hello")
	stringStack.push("world")

	stringStack.pop()
}
