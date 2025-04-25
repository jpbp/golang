package main

import "fmt"



func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	fmt.Println(actualCostInPennies,maxMessagesToSend,balance)
	for balance >= 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
		fmt.Println(actualCostInPennies,maxMessagesToSend,balance)
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}


func main() {
	fmt.Println(getMaxMessagesToSend(1.3, 10))
}
