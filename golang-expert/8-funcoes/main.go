package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(51,2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(valor)
}


func sum (a , b int) (int,error){
	if a + b >= 50{
		return 0, errors.New("soma maior que 50")
	}
	return a + b,nil
}