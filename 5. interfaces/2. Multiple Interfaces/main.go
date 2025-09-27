/*

🔑 What is Multiple Interface in Go?

In Go:
	- A struct can implement many interfaces at the same time.
	- You don’t need to write implements — just give the struct the required methods, and Go automatically knows.

*/

package main

import "fmt"

// first interface
type Reader interface {
	read() string
}

// second interface
type Writer interface {
	write(data string)
}

type File struct {
	content string
}

func (f File) read() string {
	return f.content
}

// this change the original content of the file
func (f *File) write(data string) {
	f.content = data
}

func main() {
	var r Reader
	var w Writer

	f := &File{}

	// File implements both Reader and Writer
	r = f
	w = f

	w.write("Hello Interfaces!")
	fmt.Println(r.read()) // "Hello Interfaces!"
}
