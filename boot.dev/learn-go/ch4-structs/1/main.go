package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message string
}


func main(){
	message1 := messageToSend{
		phoneNumber: 148255510981,
		message: "Thanks for signing up",
	}
	fmt.Printf("Sending message: '%s' to: %v\n",message1.message,message1.phoneNumber)

	message2 := messageToSend{
		phoneNumber: 148255510982,
		message: "Love to have you aboard!'",
	}
	fmt.Printf("Sending message: '%s' to: %v\n",message2.message,message2.phoneNumber)
}