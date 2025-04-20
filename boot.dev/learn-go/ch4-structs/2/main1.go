package main

import "fmt"

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
	
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0{
		return false
	}
	if mToSend.sender.name == "" || mToSend.sender.number == 0{
		return false
	}

	return true
}

func main(){
	user1 := user{
		name: "jp",
		number: 999290173,
	}
	user2 := user{
		name: "ana",
		number: 988580906,
	}
	message1 := messageToSend{
		message: "Oi meu amor! Bom dia, feliz pascoa" ,
		recipient: user1,
		sender: user2,
	}

	fmt.Println(canSendMessage(message1))

}