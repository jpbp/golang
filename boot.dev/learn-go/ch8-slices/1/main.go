package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messagesSends :=[3]string{primary,secondary,tertiary}
	var costSends [3]int
	cost:=0
	for i:=0; i<len(messagesSends); i++{
		cost +=  len(messagesSends[i])
		costSends[i]=cost
	}
	return messagesSends,costSends
}

func main(){
	var myInts []int
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	m,c:=getMessageWithRetries("hello","wordl","!")
	fmt.Println(m)
	fmt.Println(c)
	fmt.Println(len(myInts))
}
