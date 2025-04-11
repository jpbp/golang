package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("ta indo...")


	//Tipos de ints
	//int8,int16,int32,int64
	//Apenas o int, ele usa a arquitetura do seu computador. 
	var numero int = 100000000000
	fmt.Println(numero)


	//unit unsygned int, numero só pode ser positivo

	var numero2 uint = 3000
	fmt.Println(numero2)


	//INT32 = RUNE
	//INT8 = BYTE

	var numero3 rune = 123422
	var numero4 byte = 12

	fmt.Println(numero3,numero4)


	var numeroreal1 float32 = 123.45
	var numeroreal2 float64 = 123333333333333333.45

	fmt.Println(numeroreal1,numeroreal2)

	numeroreal3 := 1123.2323
	fmt.Println(numeroreal3)

	//strings e chars

	var str string = "texto1"
	str2 := "texto2"

	fmt.Println(str,str2)

	char := 'B'
	fmt.Println(char)

	// fim string

	//valor zero/inicial
	var num int16
	var texto string

	fmt.Println(num,texto)


	var booleano1 bool = true
	var booleano2 bool 
	fmt.Println(booleano1,booleano2)

	var erro error = errors.New("Error interno")
	fmt.Println(erro)

}
