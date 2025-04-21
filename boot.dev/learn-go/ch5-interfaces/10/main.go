package main

import "fmt"


func getExpenseReport(e expense) (string, float64) {
	email1,isEmail := e.(email)
	sms1,isSms := e.(sms)

	if isEmail{
		return email1.toAddress,e.cost()
	}else if isSms{
		return sms1.toPhoneNumber,e.cost()
	}else{
		return "",e.cost()
	}
}


func getType(e expense) {
	_,isEmail := e.(email)
	_,isSms := e.(sms)
	//_,isInvalid := e.(invalid)
	
	if isEmail{
		fmt.Println("é um email",e)
	}else if isSms{
		fmt.Println("é um sms",e)
	}else{
		fmt.Println("invalido",e)
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
	getType(email1)
	getType(sms1)
	getType(invalid1)

	fmt.Println(getExpenseReport(email1))
	fmt.Println(getExpenseReport(sms1))
	fmt.Println(getExpenseReport(invalid1))
}