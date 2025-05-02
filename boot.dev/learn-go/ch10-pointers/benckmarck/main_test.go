package main

import "testing"

func BenchmarkProcessarClientePorValor(b *testing.B) {
	cliente := Cliente{
		Nome:     "João",
		Email:    "joao@email.com",
		Endereco: "Rua Exemplo",
	}
	for i := 0; i < b.N; i++ {
		ProcessarClientePorValor(cliente)
	}
}

func BenchmarkProcessarClientePorPonteiro(b *testing.B) {
	cliente := Cliente{
		Nome:     "João",
		Email:    "joao@email.com",
		Endereco: "Rua Exemplo",
	}
	for i := 0; i < b.N; i++ {
		ProcessarClientePorPonteiro(&cliente)
	}
}
