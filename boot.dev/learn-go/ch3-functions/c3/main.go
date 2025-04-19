/*
Closure em Go

Em Go, uma função se torna um *closure* quando é definida dentro de outra função
e captura variáveis do escopo externo. Isso permite que a função "lembre"
do ambiente onde foi criada, mantendo o estado dessas variáveis mesmo após
a execução da função externa.

O compilador do Go garante que as variáveis capturadas permaneçam alocadas
na memória (heap) enquanto forem acessadas pelo closure, permitindo que o
estado seja mantido entre chamadas.

Exemplo clássico de closure:

func contador() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

Nesse exemplo, a função interna retorna uma nova função que incrementa `x`.
Cada chamada à função retornada continuará usando e atualizando a mesma
variável `x`, mesmo após `contador()` já ter terminado sua execução.

Closures são úteis para encapsular estado, criar funções geradoras,
acumuladores e construir lógica mais flexível e modular.

Regras para um closure:
1. A função é definida dentro de outra função.
2. A função usa variáveis do escopo externo.
3. A função é retornada ou mantida viva fora do seu escopo original.
*/


package main
import "fmt"

func adder() func(int) int {
	sum := 0
	return func (s int ) int {
		sum += s
		return sum
	}
}

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive

	somaDeNumeros := adder()
	fmt.Println(somaDeNumeros(1))
	fmt.Println(somaDeNumeros(2))
	fmt.Println(somaDeNumeros(3))
	fmt.Println(somaDeNumeros(4))
}