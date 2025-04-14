package main

import "fmt"

//Constante! 
//const a = "Hello, World!" 

//Variaveis de espoco global! Todos tem aceso. 
var (
	b bool = true
	c int = 10
	d string = "JP"
	e float64 = 1.222235
)

//Go é fortemente tipado, o tipo não é alterado
// Se eu quiser mudar o tipo da variavel b para int nao tem como

//Todas as variaves tem que ser usado em scope local. 
func main(){
	//var a string = "X"
	// Só pode ser usado pela primeira vez ":="
	a := "X"
	a = "AS"
	fmt.Println(a,b,c,d,e)
}

func x (){

}