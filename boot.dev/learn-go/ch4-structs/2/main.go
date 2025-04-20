package main

import "fmt"

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

func formatPrint( oneCar car){
	fmt.Printf("Meu carro é um %s, modelo %s, ele possui %v portas e sua kilometragem é de %v",oneCar.brand,oneCar.model,oneCar.doors,oneCar.mileage)
}

func main1() {
	mywheel := wheel{
		radius: 5,
		material: "rubber",
	}
	myCar := car{
		brand: "onix",
		model: "lt 1.4 aut",
		doors: 4,
		mileage: 135000,
		frontWheel: mywheel,
		backWheel: mywheel,
	}
	formatPrint(myCar)
}
