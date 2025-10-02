package main

import "fmt"

/*

Enums in Go are not like in languages such as Java, C#, or Rust.
Go doesn’t have a true enum keyword — instead, it uses a pattern with constants + iota to achieve the same behavior.

*/

/*

What is an Enum?

An enum (enumeration) is a way to define a set of named constants that represent fixed values.

Example in other languages:
enum Direction { NORTH, SOUTH, EAST, WEST }
In Go, we simulate enums with constants + iota.

--------------------------------------------------------

Enums in Go (with const and iota)
*/

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}

func main() {
	fmt.Println(North, South, East, West) // 0 1 2 3
}

/*
✨ Explanation:

iota starts at 0 and increments by 1 inside a const block.

North = 0, South = 1, etc.

Direction is a custom type (int under the hood).

*/
