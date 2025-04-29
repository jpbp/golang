package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i , msg := range messages {
		messages[i].tags = tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	m := strings.ToLower(msg.content)
	if strings.Contains(m, "urgent") {
		tags = append(tags, "Urgent")
		//return tags
	} else if strings.Contains(m, "sale") {
		tags = append(tags, "Promo")
		//return tags
	}
	//tags = append(tags, "")
	return tags
}

func main() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		{id: "003",content: "teste 123"},
		{},
		// Additional messages...
	}

	//fmt.Println(tagger(messages[3]))
	returnMessages := tagMessages(messages, tagger)

	for _, msg := range returnMessages {
		fmt.Println(msg)
	}
}
