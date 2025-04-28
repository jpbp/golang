package main

import "fmt"

func formatterHello (s string) string{
	return "Hello! "+s
}
func formatterDot (s string) string{
	return "Hello! "+s+"..."
}


func getFormattedMessages(messages []string, formatter  func(string) string) ([]string) {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
func main(){
	msg := [1]string{"Love your hair!"}
	fmt.Println(getFormattedMessages(msg[:],formatterDot))
}