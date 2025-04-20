package main

import "fmt"

type car struct {
	brand string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	bedSize int
}


func main1() {
	lanesTruck := truck{
		bedSize: 10,
		car: car{
		  brand: "Toyota",
		  model: "Camry",
		},
	  }
	  //neste caso nao precisei passar lanesTruck.car.brand, o truck imcorporou o objeto car para si, podendo acessar ele diretamente
	  fmt.Println(lanesTruck.brand) // Toyota
	  fmt.Println(lanesTruck.model) // Camry 
}