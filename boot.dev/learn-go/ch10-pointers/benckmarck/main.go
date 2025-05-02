package main



type Cliente struct {
	Nome     string
	Email    string
	Endereco string
	Compras  [1000]float64 // Array grande para simular uma struct pesada
}
var result string 
func ProcessarClientePorValor(c Cliente) {
	
	result = c.Nome // força uso

}

func ProcessarClientePorPonteiro(c *Cliente) {
	result = c.Nome

}
