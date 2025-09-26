package main

import "fmt"

// ===========================================

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backwheel  wheel
}

type wheel struct {
	radius   int
	material string
}

// ================================================

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

func main() {
	// by defalt the value store in the struct is 0
	myCar := car{}

	u1 := user{
		name:   "Abhi",
		number: 0,
	}

	u2 := user{
		name:   "Friend",
		number: 0,
	}

	msg := messageToSend{
		message:   "Hello",
		sender:    u1,
		recipient: u2,
	}

	canSendMessageResponse := canSendMessage(msg)

	fmt.Println(canSendMessageResponse)
	fmt.Println(myCar)

}
