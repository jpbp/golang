package main

import "fmt"



func getExpenseReport(e expense) (string, float64) {
	switch obj := e.(type){
		case email:		
			return obj.toAddress,e.cost() 
		case sms:
			return obj.toPhoneNumber,e.cost()
		default:
			return "",e.cost()
	}
}

func getType(e expense) () {
	switch obj := e.(type){
		case sms:
			fmt.Printf("%T\n", obj)
		case email:		
			fmt.Printf("%T\n", obj)
		default:
			fmt.Printf("%T\n", obj)
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main(){
	email1 := email{isSubscribed: false,body: "testeEmail",toAddress: "40028922"}
	sms1 := sms{isSubscribed: true,body: "teste123",toPhoneNumber: "999290173"}
	invalid1 := invalid{}

	getExpenseReport(email1)
	getExpenseReport(sms1)
	getExpenseReport(invalid1)
}