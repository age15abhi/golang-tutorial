/* =================================================

🟢 What is an Empty Struct?

An empty struct in Go is a struct with no fields:

		|	struct{} |

It takes zero bytes of memory.

Often used as a signal or placeholder rather than storing data.

✅ Key point: Even though it’s a struct, it doesn’t occupy memory, so it’s extremely lightweight.
================================================== */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct{}
	fmt.Println(x)                // {}
	fmt.Println(unsafe.Sizeof(x)) // 0
}
