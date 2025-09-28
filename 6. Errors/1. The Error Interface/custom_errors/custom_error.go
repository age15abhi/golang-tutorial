package main

import "fmt"

/*

✅ Example 2: Custom error type

You can create your own error with extra details

*/

type DivideError struct {
	Dividend int
	Divisor  int
	message  string
}

func (e DivideError) Error() string {
	return fmt.Sprintf("Divide error: %d / %d - %s", e.Dividend, e.Divisor, e.message)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideError{a, b, "division by zero is not allowed"}
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
