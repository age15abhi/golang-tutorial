/*

1️⃣ What is a Slice?

A slice is a flexible, dynamically-sized view into an array.

Unlike arrays, slices don’t have a fixed length.

A slice is made of three things internally:
	- Pointer → points to the underlying array
	- Length → number of elements in the slice
	- Capacity → size of the underlying array starting from the first element of the slice

Slices are more powerful than arrays because they can grow and shrink dynamically.

*/

package main

import "fmt"

func main() {
	// --------------- create slice -----------------
	// from an array
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4] // slice from index 1 to 3
	fmt.Println(s)

	// Direct slice literal
	directSliceLiteral := []int{1, 2, 3, 4, 5}
	fmt.Println(directSliceLiteral)

	// Using make
	sliceLiteralUsingmake1 := make([]int, 5) // length = 5, capacity = 5
	fmt.Println(sliceLiteralUsingmake1)
	sliceLiteralUsingmake2 := make([]int, 3, 10) // length = 3, capacity = 10
	fmt.Println(sliceLiteralUsingmake2)
	fmt.Println(len(sliceLiteralUsingmake2), cap(sliceLiteralUsingmake2))

	// --------------- Slice Operations -----------------
	// 1. Accessing element
	sliceForOperation := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceForOperation[0])
	sliceForOperation[1] = 50
	fmt.Println(sliceForOperation)

	// 2. Slicing a slice
	sliceForOperation1 := []int{10, 20, 30, 40, 50}
	sfo1 := sliceForOperation1[1:4]
	sfo2 := sliceForOperation1[:3]
	sfo3 := sliceForOperation1[2:]
	fmt.Println(sfo1)
	fmt.Println(sfo2)
	fmt.Println(sfo3)

	// 3. Appending elements
	sliceForOperation2 := []int{10, 20, 30, 40, 50}
	sliceForOperation2 = append(sliceForOperation2, 60)
	sliceForOperation2 = append(sliceForOperation2, 70, 80, 90)
	fmt.Println(sliceForOperation2)

	// 4. Copying slices
	sliceForOperation3 := []int{1, 2, 3, 4, 5}
	sliceForOperation4 := make([]int, len(sliceForOperation3))
	copy(sliceForOperation4, sliceForOperation3)
	fmt.Println(sliceForOperation4)

	// --------------- Multi-dimensional slices ---------------
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix[0][1]) // 2

	// --------------- Nil slice ---------------
	var nillSlice []int
	fmt.Println(nillSlice == nil)
	fmt.Println(len(nillSlice), cap(nillSlice)) // 0 0
	// we can append the value to the slices
	nillSlice = append(nillSlice, 1, 2, 3)
	fmt.Println(nillSlice)

	// ------- Slice Internals (Memory sharing) -------
	memorySharingArray := [5]int{1, 2, 3, 4, 5}
	sliceForOperation5 := memorySharingArray[1:4] // 1-3
	sliceForOperation6 := sliceForOperation5[1:3] // 1-2

	sliceForOperation6[0] = 99
	fmt.Println(sliceForOperation6)
	fmt.Println(memorySharingArray)

	// --------------- Slice tricks --------------------
	sliceForOperation7 := []int{1, 2, 3, 4, 5}

	// check if slice contains an element
	found := false

	for _, value := range sliceForOperation7 {
		if value == 3 {
			found = true
			break
		}
	}

	fmt.Println(found)

	// Remove element by index
	sliceForOperation8 := []int{1, 2, 3, 4, 5}
	i := 2
	sliceForOperation8 = append(sliceForOperation8[:i], sliceForOperation8[i+1:]...)
	fmt.Println(sliceForOperation8)

	// [1 , 2] --> take all element after 3  --> [4 , 5]
	// append [1,2,4,5]
}
