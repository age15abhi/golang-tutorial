/*

In Go, errors are values.
Instead of exceptions (like Java, Python, JS), Go uses an error interface to represent and handle errors.

The built-in error interface is defined as:

type error interface {
    Error() string
}

So, any type that implements the Error() string method becomes an error.


-------------------------------------------------------

✅ Why use the error interface?

Consistent error handling
All functions in Go return errors in a standard way: value, error.

Explicit error handling
Forces the programmer to deal with errors rather than ignoring them.

Custom error types
You can define your own error struct with additional context.

Simple and flexible
Errors are just values — can be returned, stored, or wrapped.

*/

package main

import (
	"errors"
	"fmt"
)

// ✅ Example 1: Built-in error

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
