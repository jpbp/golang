package main

import (
	"fmt"
)

func main() {


	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
}


func sum (numeros ...int) (int){
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}