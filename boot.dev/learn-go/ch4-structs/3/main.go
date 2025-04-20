package main

import "fmt"

type car struct {
	brand   string
	model   string
	doors   int
	mileage int
	// wheel is a field containing an anonymous struct
	wheel struct {
		radius   int
		material string
	}
}

func main() {
	myCar := struct {
		brand string
		model string
	}{
		brand: "Toyota",
		model: "Camry",
	}
	fmt.Println(myCar)
}
