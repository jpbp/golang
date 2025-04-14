package main

import "fmt"

func main() {
	
	var meuArray [4] int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40
	fmt.Println("Tamanho do array: ",len(meuArray)-1)
	fmt.Println("Meu array: ",meuArray)
	fmt.Println("Ultimo Elemento: ",meuArray[len(meuArray)-1])

	for i,v := range meuArray{
		fmt.Printf("O valor do indice %d é %d\n",i,v)
	}
}
