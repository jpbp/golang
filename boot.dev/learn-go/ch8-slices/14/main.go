package main

import "fmt"

func main() {
	fmt.Println("-----------------------EXEMPLO 1 ----------------")
	a := make([]int, 3)
	fmt.Println("len of a:", len(a))
	fmt.Println("cap of a:", cap(a))
	// len of a: 3
	// cap of a: 3

	b := append(a, 4)
	fmt.Println("appending 4 to b from a")
	fmt.Println("b:", b)
	fmt.Println("addr of b:", &b[0])
	// appending 4 to b from a
	// b: [0 0 0 4]
	// addr of b: 0x44a0c0

	c := append(a, 5)
	fmt.Println("appending 5 to c from a")
	fmt.Println("addr of c:", &c[0])
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	// appending 5 to c from a
	// addr of c: 0x44a180
	// a: [0 0 0]
	// b: [0 0 0 4]
	// c: [0 0 0 5]

	fmt.Println("-----------------------EXEMPLO 2 ----------------")

	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	fmt.Println("cap of i:", cap(i))
	// len of i: 3
	// cap of i: 8

	j := append(i, 4)
	fmt.Println("appending 4 to j from i")
	fmt.Println("j:", j)
	fmt.Println("addr of j:", &j[0])
	// appending 4 to j from i
	// j: [0 0 0 4]
	// addr of j: 0x454000

	g := append(i, 5)
	fmt.Println("appending 5 to g from i")
	fmt.Println("addr of g:", &g[0])
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("g:", g)
	// appending 5 to g from i
	// addr of g: 0x454000
	// i: [0 0 0]
	// j: [0 0 0 5]
	// g: [0 0 0 5]
}
