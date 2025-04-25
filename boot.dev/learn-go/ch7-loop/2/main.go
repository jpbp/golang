package main

import "fmt"

func maxMessages(thresh int) int {
	totalCost := 0
	qtdMessages := 0
	for i:=0; totalCost<=thresh; i++{
		totalCost += 100 + i 
		qtdMessages=i
	}
	return qtdMessages
}
func main(){
	fmt.Println(maxMessages(200))
}