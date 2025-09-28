/*

A map in Go is a key-value store (similar to dictionaries in Python, objects in JS, or hashmaps in Java).
They’re implemented as hash tables under the hood.

*/

package main

import "fmt"

type Person struct {
	Age  int
	City string
}

func main() {

	// declare and empty map
	// var m map[string]int

	// declare the map using the make function
	myMap := make(map[string]int)

	// add key-value pair
	myMap["alice"] = 25
	myMap["abhi"] = 20

	// map with shortand initilization
	map1 := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// -------------- Accessing Values ------------------
	fmt.Println(myMap["alice"])
	fmt.Println(myMap["abhi"])
	fmt.Println(map1["Bob"])

	// ------------------- Check if key exist -------------
	age, exists := map1["kaliya"]

	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("Bob not found")
	}

	// Updating map value
	map1["Alice"] = 30 // updates Alice’s age

	fmt.Println(map1)

	// Delete the key from the map
	delete(map1, "Bob")
	fmt.Println(map1)

	// Iterate over a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Length of a map
	fmt.Println(len(myMap))

	//  Maps of Structs
	person := make(map[string]Person)

	person["Alice"] = Person{Age: 20, City: "delhi"}
	person["Abhi"] = Person{Age: 21, City: "rewari"}

	fmt.Println(person)

	// Nested maps
	grades := map[string]map[string]int{
		"Alice": {
			"Math":    90,
			"Science": 85,
		},
		"Bob": {
			"Math":    70,
			"Science": 60,
		},
	}

	fmt.Println(grades["Alice"]["Math"]) // 90

}
