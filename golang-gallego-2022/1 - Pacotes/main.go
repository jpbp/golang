package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("teste@teste.com.br")
	fmt.Println(erro)

}