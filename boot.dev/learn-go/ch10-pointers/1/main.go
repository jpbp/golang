package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func main(){
	m1:= Message{Recipient: "Recipient: Pumpkin",Text: "I love you, Honey Bunny."}
	fmt.Println(getMessageText(m1))
}