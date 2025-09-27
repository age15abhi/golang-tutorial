package main

import "fmt"

// ==============================================

type car struct {
	brand   string
	model   string
	doors   int
	mileage int
	wheel   struct {
		radius   int
		material string
	}
}

// ==============================================

func marksAverage(student struct {
	name  string
	marks [3]int
}) {

	// find the sum of the mars
	sum := 0

	for _, m := range student.marks {
		sum += m
	}

	avg := float64(sum) / float64(len(student.marks))

	fmt.Printf("Student %s has average marks %.2f\n", student.name, avg)
}

func main() {

	// -----------------------------------------------
	book := struct {
		title  string
		author string
		price  float64
	}{
		title:  "effective go",
		author: "abhishek",
		price:  99.9,
	}

	// -----------------------------------------------

	student := struct {
		name  string
		marks [3]int
	}{
		name:  "abhishek",
		marks: [3]int{90, 88, 90},
	}

	// -----------------------------------------------
	/* Nested named and the anonymous  */

	myCar := car{
		brand:   "audi",
		model:   "v12",
		doors:   4,
		mileage: 54,
		wheel: struct {
			radius   int
			material string
		}{
			radius:   20,
			material: "iron",
		},
	}

	// ------------------------------------------------

	// call  the marks average function
	marksAverage(student)

	// here printing the book detail
	fmt.Println(book)
	fmt.Println(myCar)

}
