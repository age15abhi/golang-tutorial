package main

/*

🔹 Real-World Use Cases of Enums in Go

*/

// 1. State machine
// Use case: Process lifecycle, task states.
type state int

const (
	idle    state = iota // 0
	Running              // 1
	Stopped              //2
)

// 2.HTTP Methods
type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

// 3.Days of the Week
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
)

// Permissions (Bit Flags) → Combine multiple flags easily.

// Game Development → Player states, weapon types, levels.
