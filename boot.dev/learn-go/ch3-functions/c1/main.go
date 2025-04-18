package main

import "fmt"

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	Calculetecost := calculateBaseBill(costPerMessage,numMessages)
	finalDiscount := calculateDiscountRate (numMessages) 
	finalCost := Calculetecost - Calculetecost * finalDiscount 
	return finalCost
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000  {
		return 0.20
	}else if messagesSent > 1000{
		return 0.10
	}
	return 0.0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main (){
	fmt.Println(calculateFinalBill(1,500))
	fmt.Println(calculateFinalBill(1,1000))
	fmt.Println(calculateFinalBill(1,1001))
	fmt.Println(calculateFinalBill(1,5000))
	fmt.Println(calculateFinalBill(1,5001))
}