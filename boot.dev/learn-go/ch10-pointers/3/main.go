package main

import "fmt"

func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

func (analytics Analytics) toFormmat(){
	fmt.Printf("Messagens Total: %v, Messagens Failed: %v, Messages Succeeded: %v\n",analytics.MessagesSucceeded,analytics.MessagesFailed,analytics.MessagesSucceeded)
}
type Message struct {
	Recipient string
	Success   bool
}

func getMessageText (analytics *Analytics, NewMessages []Message ){
	
	for _,message := range NewMessages{
		if (message.Success){
			(*analytics).MessagesSucceeded++
		}else{
			(*analytics).MessagesFailed++
		}
		(*analytics).MessagesTotal++
	}
	
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    
	analytics := Analytics{MessagesTotal: 0,MessagesFailed: 0,MessagesSucceeded: 0}
	message1 := Message{Recipient: "teste1",Success: true}
	message2 := Message{Recipient: "teste12",Success: false}
	message3 := Message{Recipient: "teste123",Success: true}
	message4 := Message{Recipient: "teste1234",Success: false}
	message5 := Message{Recipient: "teste12345",Success: true}

	//getMessageText(&analytics,message1)
	//analytics.toFormmat()

	//getMessageText(&analytics,message2)
	//analytics.toFormmat()

	//getMessageText(&analytics,message3)
	//analytics.toFormmat()

	//getMessageText(&analytics,message4)
	//analytics.toFormmat()

	//getMessageText(&analytics,message5)
	//analytics.toFormmat()


	msgs := []Message{message1,message2,message3,message4,message5}
	getMessageText(&analytics,msgs)
	analytics.toFormmat()
}