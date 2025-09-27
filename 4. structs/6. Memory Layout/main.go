/* =================================================

🟢 1. What is Memory Layout in Context of Structs?

In Go, a struct is a contiguous block of memory that stores its fields.

Each field is stored sequentially in memory, but Go may add padding to align data for CPU efficiency.

The order of fields matters because it can affect struct size and performance.

================================================== */

package main

import (
	"fmt"
	"unsafe"
)

// Example:

type MyStruct struct {
	b int64 // 8 bytes
	a int8  // 1 byte
	c int8  // 1 byte
}

/* =================================================

🟢 2. Memory Management in Structs

Go is a garbage-collected language, so you don’t manually free memory like C.
But you can control memory usage:

================================================== */

// 1. Use pointers for large structs to avoid copying:

type BigStruct struct {
	data [10000]int
}

func process(s *BigStruct) { /* ... */ } // pass pointer

// 2. Reuse structs with slices/pools (sync.Pool) for high-performance apps.

// 3. Avoid unnecessary allocations:
// - Use slices instead of arrays if size is dynamic.
// - Embed structs instead of copying big structs.

func main() {
	var s MyStruct
	fmt.Println("Size of struct:", unsafe.Sizeof(s))
}
