/*

# Rules of Packages
- The folder name = package name (by convention).

- If an identifier (function, variable, struct) starts with a capital letter, it is exported (public).

- If it starts with a lowercase letter, it is unexported (private).

*/

package mathutils

func Add(a, b int) int {
	return a + b
}

// Exported
func Multiply(a, b int) int {
	return a * b
}

// Not exported
func subtract(a, b int) int {
	return a - b
}
