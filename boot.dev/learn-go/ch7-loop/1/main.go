package main

import "fmt"

func bulkSend(numMessages int) float64 {
	costMessage := 1.0
	costTax := 0.0
	costMessageTotal := 0.0
	costTaxTotal := 0.0
	for i:=0; i<numMessages; i++ {
		costMessageTotal = costMessage + costTaxTotal
		costMessage = costMessage + 1
		costTax = costTax + 0.01
		costTaxTotal = costTax + costTaxTotal 
	}
	return costMessageTotal
}

func main(){
	fmt.Println(bulkSend(0))
}