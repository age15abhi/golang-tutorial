package main

import "fmt"

//  Example 2: Anonymous Struct

func main() {
	person := struct {
		name   string
		number int
	}{
		name:   "Abhi",
		number: 123456789,
	}

	// -----------------------------------------------------
	// Anonymous Struct inside another Struct

	msg := struct {
		sender struct {
			name   string
			number int
		}
		recipient struct {
			name   string
			number int
		}
		message string
	}{
		sender: struct {
			name   string
			number int
		}{name: "abhishek", number: 123456789},
		recipient: struct {
			name   string
			number int
		}{name: "annu", number: 123456789},
		message: "Hello Friends",
	}

	// ---------------------------------------------------
	fmt.Println(person)

	// print the name and number
	fmt.Println(person.name)
	fmt.Println(person.number)

	// ---------------------------------------------------
	fmt.Println(msg)
}
