
# Benchmark: Structs por Valor vs Ponteiro em Go

Este projeto testa e documenta a diferença de performance entre passar uma struct grande **por valor** e **por ponteiro** em Go.

---

## 🔍 Objetivo

Entender em quais situações o uso de ponteiros traz ganhos de performance reais em comparação à passagem de structs por valor.

---

## 📦 Struct utilizada no benchmark

```go
type Cliente struct {
    Nome     string
    Email    string
    Endereco string
    Compras  [1000]float64
}
```

Essa struct contém um array grande para simular um custo realista de cópia.

---

## 🔁 Funções testadas

```go
// Global para evitar otimização do compilador
var result string

// Função que recebe a struct por valor
func ProcessarClientePorValor(c Cliente) {
    result = c.Nome
}

// Função que recebe a struct por ponteiro
func ProcessarClientePorPonteiro(c *Cliente) {
    result = c.Nome
}
```

---

## 🧪 Execução do benchmark

Comando:

```bash
go test -bench=. -benchmem
```

Resultado:

```
BenchmarkProcessarClientePorValor-8      19071240     64.70 ns/op
BenchmarkProcessarClientePorPonteiro-8  1000000000     0.6906 ns/op
```

---

## ✅ Conclusão

| Estratégia            | Tempo médio   | Interpretação                        |
|-----------------------|---------------|--------------------------------------|
| Passagem por valor    | ~64.70 ns/op  | 🚨 Mais lenta (cópia da struct)      |
| Passagem por ponteiro | ~0.69 ns/op   | ⚡ Muito mais rápida (sem cópia)     |

- Usar ponteiros pode trazer ganhos de performance **significativos**, especialmente com structs grandes.
- Passar structs por valor deve ser evitado em cenários de alta carga ou com uso repetitivo.

---

## 💡 Boas práticas

Use **ponteiros** quando:
- A struct for grande.
- Você deseja modificar o dado original.
- Você quer evitar múltiplas alocações de memória.
- Você quer melhorar a performance de código repetitivo (ex: loops, jobs, APIs).

Use **valor** quando:
- Você quer uma cópia independente.
- A struct for pequena.
- Imutabilidade for desejada no contexto da função.

---

## 📁 Estrutura do projeto

```
.
├── go.mod
├── main.go           # Contém as funções benchmark
└── main_test.go      # Arquivo de benchmark
```

