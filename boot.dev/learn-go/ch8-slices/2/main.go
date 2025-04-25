package main


import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro"{
		return messages[:],nil
	}else if plan == "free"{
		return messages[:2],nil
	}else{
		return nil,errors.New("unsupported plan")
	}
}


func main() {
	/*
	arrayname[lowIndex:highIndex]
	arrayname[lowIndex:]
	arrayname[:highIndex]
	arrayname[:]
	*/
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[:2])
	fmt.Println(primes[0:])
	fmt.Println(primes[:5])
	fmt.Println(primes[:])
		//var myInts [10]int
}
