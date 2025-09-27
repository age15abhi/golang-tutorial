package main
import "fmt"

type messageToSend struct {
	phoneNumber		int
	message 		string
}

func main () {

	msg := messageToSend {
		phoneNumber: 123456789,
		message: "hello abhishek",
	}


	fmt.Println(msg)

}

