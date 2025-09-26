package main

import "fmt"

/*

Create a struct Engine with fields:
	horsepower (int)
	fuelType (string)

*/

type Engine struct {
	horsepower int
	fuelType   string
}

/*
Create another struct Car that embeds Engine and also has:

	brand (string)
	model (string)
*/

type Car struct {
	Engine
	brand string
	model string
}

func main() {
	car := Car{
		Engine: Engine{horsepower: 750, fuelType: "petrol"},
		brand:  "audi",
		model:  "v12",
	}

	fmt.Println(car.horsepower)
}
