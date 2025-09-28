/*

🔹 What is panic in Go?
	- panic is Go’s way of saying “something went really wrong, I can’t continue execution safely”.
	- It immediately stops the normal execution of the program, unwinds the call stack, and if not recovered, the program crashes.

In contrast:

	- error is for expected problems that the caller can handle (like division by zero, file not found, invalid input).

	- panic is for unexpected or unrecoverable problems (like array out of bounds, nil pointer dereference, corrupt memory).


-----------------------------------------------------------

🔹 How does panic work?

1. When you call panic("message"), Go:

	- Prints the panic message.
	- Prints the stack trace (which functions were running).
	- Unwinds the stack and stops the program (unless recovered).

2. You can recover from a panic using recover() inside a defer function.
*/

package main

import "fmt"

func divideError(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func dividePanic(a, b int) (result int) {

	// here we are adding the defer revocer() function so we can recover from the crash the program

	defer func() {
		if r := recover(); r != nil {
			// recover() catches the panic value
			fmt.Println("Recover from panic:", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("division by zero caused panic!") // unexpected situation
	}
	return a / b
}

func main() {
	// ✅ Error usage
	if res, err := divideError(10, 0); err != nil {
		fmt.Println("Handled with error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	// ❌ Panic usage (will crash program if not recovered)
	fmt.Println("Now calling dividePanic...")
	res := dividePanic(10, 0)
	fmt.Println("This line will never run:", res)
}
