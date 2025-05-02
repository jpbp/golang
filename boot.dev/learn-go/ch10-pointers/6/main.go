package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil{
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}


func main (){
	var sp * string = nil
	removeProfanity(sp)
	fmt.Println(sp)
}