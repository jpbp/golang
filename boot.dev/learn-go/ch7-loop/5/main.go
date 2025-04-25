package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i:=1; i<=max; i++{
		countPrimes:=0
		for j:=1; j<=i; j++{
			if(i%j==0){
				countPrimes++
			}
		}
		if(countPrimes==2){
			fmt.Println(i)
		}
		countPrimes=0
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
