package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	msg := []Message{}
	for _,m := range messages{
		if (m.Type()==filterType){
			msg = append(msg, m)
		}	
	}
	return msg
}

func main(){
	t1 := TextMessage{Sender: "jp1",Content: "teste11"}
	t2 := TextMessage{Sender: "jp2",Content: "teste12"}
	m1 := MediaMessage{Sender: "jp3",MediaType: "mp4",Content: "teste21"}
	m2 := MediaMessage{Sender: "jp4",MediaType: "mp4",Content: "teste22"}
	l1 := LinkMessage{Sender: "jp5",URL: "drive.com",Content: "teste31"}
	l2 := LinkMessage{Sender: "jp6",URL: "drive.com",Content: "teste32"}

	msg := []Message{}
	msg = append(msg,l1)
	msg = append(msg,m1)
	msg = append(msg,t1)
	msg = append(msg,t2)
	msg = append(msg,m2)
	msg = append(msg,l2)

	fmt.Println(filterMessages(msg,"teste"))

}
