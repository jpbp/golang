package main

import (
	"errors"
	"fmt"
)

func main() {
	valor := sum(51,2)
	fmt.Println(valor)
}


func sum (a , b int) (int){
	return a + b
}