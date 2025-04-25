package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))

	for i:=0; i <len(messages); i++{
		cost := float64(len(messages[i])) * 0.01
		messagesCost[i] = cost
	}
	return messagesCost[:]
}

func main(){
	s := [3]string{"jp","teste","foi"}
	c := getMessageCosts(s[:])
	fmt.Println(c)
	
}
