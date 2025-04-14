package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	salarios := map[string]int{
		"João": 1500,
		"Maria": 2000,
		"José": 2500,}
	
	for nome, salario := range salarios {
		fmt.Printf("O nome da pessoa é %s e o salário é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de: %d\n", salario)
	}
	//delete(salarios,"João")
	//fmt.Println(salarios)

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["jp"]=1500
	
	
}

