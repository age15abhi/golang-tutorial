package main

import (
	"fmt"
	"myproject/mathutils"   // import local package
	"myproject/stringutils" // import local package

	// Here come the third party package
	"github.com/google/uuid"
)

func main() {
	sum := mathutils.Add(5, 7)
	text := stringutils.Concat("Go", "Lang")

	fmt.Println("Sum:", sum)     // 12
	fmt.Println("Concat:", text) // GoLang

	id := uuid.New() // Generate a new UUID
	fmt.Println("UUID:", id)
}
