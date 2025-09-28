package main

import "fmt"

type Desciber interface {
	Descibe() string
}

type CalculatorError struct {
	Operation string
	Operand1  int
	Operand2  int
	Message   string
}

// ---------------- Implement Interfaces ----------------
func (e CalculatorError) Error() string {
	return fmt.Sprintf("Error in %s: %d and %d -> %s",
		e.Operation, e.Operand1, e.Operand2, e.Message)
}

func (e CalculatorError) Descibe() string {
	if e.Message == "" {
		return "Everything is fine"
	}
	return fmt.Sprintf("Detailed description: operation %s failed because %s (operands: %d, %d)",
		e.Operation, e.Message, e.Operand1, e.Operand2)
}

// ---------------- Functions ----------------
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CalculatorError{Operation: "divide",
			Operand1: a,
			Operand2: b,
			Message:  "division by zero is not allowed",
		}
	}

	return a / b, nil
}

func multiply(a, b int) (int, Desciber) {
	if a < 0 || b < 0 {
		return 0, CalculatorError{
			Operation: "Multiply",
			Operand1:  a,
			Operand2:  b,
			Message:   "negative numbers are not allowed",
		}
	}

	return a * b, CalculatorError{}
}

// ---------------- Main ----------------
func main() {

	// ------------- divide example --------------
	result, err := divide(10, 0)

	if err != nil {
		// type assertion
		if calcErr, ok := err.(CalculatorError); ok {
			fmt.Println(calcErr.Error())
			fmt.Println(calcErr.Descibe())
		}
	} else {
		fmt.Println(result)
	}

	// -------------------- multiply example ---------
	result, desc := multiply(-2, 3)

	if calcErr, ok := desc.(CalculatorError); ok {

		fmt.Println(calcErr.Error())
		fmt.Println(calcErr.Descibe())
	} else {
		fmt.Println(result)
	}
}
