package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0,errors.New("no dividing by 0")
	}
	return x / y, nil
}

func main(){
	fmt.Println(divide(15,15))
	fmt.Println(divide(15,0))
}