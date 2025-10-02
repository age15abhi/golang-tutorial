package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt(prompt string) int64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator. Goodbye!")
			os.Exit(0)
		}

		n, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}
		return n
	}
}

func ReadOperation(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "+" || input == "-" || input == "*" || input == "/" {
			return input
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator. Goodbye!")
			os.Exit(0)
		}
		fmt.Println("Invalid operation, try again")
	}
}

func Calculate(number1, number2 int64, operation string) {

	switch operation {
	case "+":
		fmt.Println("Result:", number1+number2)
	case "-":
		fmt.Println("Result:", number1-number2)
	case "*":
		fmt.Println("Result:", number1*number2)
	case "/":
		if number2 != 0 {
			fmt.Println("Result:", number1/number2)
		} else {
			fmt.Println("Cannot divide by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
}

func main() {
	for {

		number1 := ReadInt("Enter number 1 : ")
		number2 := ReadInt("Enter number 2 : ")

		operation := ReadOperation("Enter opeation ?\n")
		Calculate(number1, number2, operation)

	}
}
