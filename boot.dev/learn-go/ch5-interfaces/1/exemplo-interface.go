package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}
func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main1(){
	c1 := circle{
		radius: 10,
	}

	r1 := rect {
		width: 1,
		height: 5,
	}
	printShapeData(c1)
	printShapeData(r1)
}
