package main

import "fmt"

func main(){
	var var1 string = "variavel 1"
	var2 := "variavel 2"
	
	fmt.Println(var1)
	fmt.Println(var2)

	var (
		variavel3 string = "lá lá"
		variavel4 string = "jpjp"
	)
	fmt.Println(variavel3,variavel4)
	
	variavel5,variavel6 := "variavel 5" , "variavel6"
	fmt.Println(variavel5,variavel6)


	const constant1 string = "Constante 1"
	fmt.Println(constant1)


	//alterando duas variaveis
	variavel5, variavel6 = variavel6,variavel5
	fmt.Println(variavel5,variavel6)
}
