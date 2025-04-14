package main

import "fmt"

// Constante!
const a = "Hello, World!"

// Tipo id Criado
type ID int

// Variaveis de espoco global! Todos tem aceso.
var (
	b bool    = true
	c int     = 10
	d string  = "JP"
	e float64 = 1.222235
	f ID      = 1
)

//Go é fortemente tipado, o tipo não é alterado
// Se eu quiser mudar o tipo da variavel b para int nao tem como

// Todas as variaves tem que ser usado em scope local.
func main() {
	//var a string = "X"
	// Só pode ser usado pela primeira vez ":="
	//a := "X"
	//fmt.Println(f)
	/*
	%T tipo da variavel
	%v valor da variavel
	*/
	fmt.Printf("O tipo de E é %T",e)
	fmt.Printf("O tipo de E é %v",e)
	fmt.Printf("O tipo de F é %T",f)
}
