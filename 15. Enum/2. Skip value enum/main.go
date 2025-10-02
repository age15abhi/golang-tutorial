package main

import "fmt"

// 🔹 Example 2: Skip Values
const (
	A = iota
	B
	_
	C
)

// Example 3: Bit Flags (Power of 2 Enums)
const (
	Read    = 1 << iota // 1
	Write               // 2
	Execute             // 4
)

// 🔹 Example 4: Custom Start Value
const (
	Sunday = iota + 1 // start from 1
	Monday
	Tuesday
)

// 🔹 Example 5: Typed vs Untyped
// Untyped constants
const (
	Red = iota
	Green
	Blue
)

// Typed enum
type Color int

const (
	RedColor Color = iota
	GreenColor
	BlueColor
)

// 🔹 Example 6: Auto-Generated Stringer
// Go has a tool to generate enum string names automatically:
// go install golang.org/x/tools/cmd/stringer@latest
//go:generate stringer -type=Direction
// go generate

func main() {
	fmt.Println(A, B, C)

	perms := Read | Write
	fmt.Println(perms)              // 3
	fmt.Println(perms&Read != 0)    // true
	fmt.Println(perms&Execute != 0) // false
}
