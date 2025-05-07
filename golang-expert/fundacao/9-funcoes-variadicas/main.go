package main

import (
	"fmt"
)

func main() {
	valor := sum(1,2,3,4,5,5,6,7,8,9,10)
	fmt.Println(valor)
}


func sum (numeros ...int) (int){
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}