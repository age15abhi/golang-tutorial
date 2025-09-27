/*

📝 Assignment: Vehicles and Speeds
Task:

1.  an interface called Vehicle with one method:

```Speed() int```

→ This method should return the speed of the vehicle.

2.Create two structs:

	- Car (with a MaxSpeed field)
	- Bike (with a MaxSpeed field)

3. Make both Car and Bike implement the Vehicle interface.

	- Their Speed() method should return their MaxSpeed.

4. Write a function:

```func PrintSpeed(v Vehicle)```


→ This should print the speed of any vehicle passed to it.

In main():

	- Create a Car (say speed 120)
	- Create a Bike (say speed 80)
	- Call PrintSpeed for both.

*/

package main

import "fmt"

// speed name function can be create many time
// by the different struct object
type Vehicle interface {
	speed() int64
}

type Car struct {
	MaxSpeed int64
}

type Bike struct {
	MaxSpeed int64
}

// here i create the speed function with the same name
// but for the different type
func (c Car) speed() int64 {
	return c.MaxSpeed
}

func (b Bike) speed() int64 {
	return b.MaxSpeed
}

func PrintSpeed(v Vehicle) {
	fmt.Println(v.speed())
}

func main() {
	c := Car{
		MaxSpeed: 120,
	}

	b := Bike{
		MaxSpeed: 80,
	}

	// i can also run the method like
	// b.speed()
	// c.speed()

	PrintSpeed(c)
	PrintSpeed(b)
}
