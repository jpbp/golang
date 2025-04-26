package main

import (
	"errors"
	"fmt"
)

type Compra struct{
	Valor float64
	Descricao string
}

type Pagamento interface{
	Pagar (Compra Compra) (float64,error)
}

type CartaoCredito struct{}
func (cartaoCredito CartaoCredito) Pagar(compra Compra ) (float64,error){
	return compra.Valor*1.05,nil
}
type Boleto struct{}

func (Boleto Boleto) Pagar(compra Compra ) (float64,error){
	return compra.Valor+2.50,nil
}
type Pix struct{}
func (pix Pix) Pagar(compra Compra ) (float64,error){
	if (compra.Valor > 10.0){
		return compra.Valor,nil
	}
	return 0.0,errors.New("É permitido apenas compras acima de 10 REAIS!!!")
}

func processarPagamento(pag Pagamento,compra Compra){
	fmt.Println(pag.Pagar(compra))
}

func main(){
	c1 := Compra{Valor: 100.0,Descricao: "Fone de Ouvido"}
	c2 := Compra{Valor: 250.0,Descricao: "Teclado Mecanico"}
	c3 := Compra{Valor: 5.0,Descricao: "Caneta azul"}
	c4 := Compra{Valor: 80.0,Descricao: "Mouse Gamer"}

	boleto := Boleto{}
	cc := CartaoCredito{}
	pix := Pix{}

	processarPagamento(cc,c1)
	processarPagamento(boleto,c2)
	processarPagamento(pix,c3)
	processarPagamento(pix,c4)
}