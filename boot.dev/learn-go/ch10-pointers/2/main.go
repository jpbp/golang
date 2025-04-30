package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")

	*message = strings.ReplaceAll(*message, "shiz", "****")

	*message = strings.ReplaceAll(*message, "shiz", "*****")

}
func main() {
	var p *int
	fmt.Printf("value of p: %v\n", p)
	myString := "hello"      // myString is just a string
	myStringPtr := &myString // myStringPtr is a pointer to myString's address

	fmt.Printf("value of myStringPtr: %v\n", myStringPtr)
	*myStringPtr = "English, motherfubber, do you speak it?" // set myString through the pointer
	removeProfanity(&myString)
	fmt.Printf("value of myString: %s\n", *myStringPtr)
	fmt.Printf("value of myString: %s\n", myString)

}
