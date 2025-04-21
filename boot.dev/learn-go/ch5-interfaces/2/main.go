package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type name interface {
	getName() string
}

type salary interface {
	getSalary() int
}

func testName(n name)string{
	return n.getName()
}

func formatPrint (em employee) {
	fmt.Printf("Funcionario: %s, salario: %v\n",em.getName(),em.getSalary())
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}
type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main(){

	ft := fullTime{
		name: "jp",
		salary: 9000,
	}
	c := contractor{
		name: "ana",
		hourlyPay: 50,
		hoursPerYear: 500,
	}

	formatPrint(ft)
	formatPrint(c)
	fmt.Println(testName(c))
	fmt.Println(testName(ft))
}