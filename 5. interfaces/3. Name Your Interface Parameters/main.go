/*

🔑 Naming Interface Parameters

When you define an interface method in Go, you can give names to the parameters (or not).

*/

package main

/*
Here, you give names (p, n, err) to make the purpose clearer.

Useful when:
- The function is complex.
- You want documentation-like clarity.
- It helps other developers understand what the parameters/returns mean.
*/

import "fmt"

// Interface with named parameters
type Logger interface {
	Log(message string, level int)
}

// Struct implementing Logger
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string, level int) {
	switch level {
	case 1:
		fmt.Println("[INFO]:", message)
	case 2:
		fmt.Println("[WARNING]:", message)
	case 3:
		fmt.Println("[ERROR]:", message)
	default:
		fmt.Println("[UNKNOWN]:", message)
	}
}

// Function that accepts Logger
func DoSomething(l Logger) {
	l.Log("Starting process", 1)
	l.Log("Something seems off", 2)
	l.Log("Process failed", 3)
}

func main() {
	logger := ConsoleLogger{}
	DoSomething(logger)
}
