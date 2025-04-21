package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed{
		return 2*len(e.body)
	}
	return 5*len(e.body)
}

func (e email) format() string {
	if e.isSubscribed{
		s := fmt.Sprintf("'%s' | Subscribed",e.body)
		return s
	}
	return fmt.Sprintf("'%s' | Not Subscribed",e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}
func printInfo(e expense, f formatter) {
	fmt.Println("Custo do email:", e.cost())
	fmt.Println("Formatação do email:", f.format())
}

type email struct {
	isSubscribed bool
	body         string
}

func main (){
	e := email{
		isSubscribed: false,
		body: "general kenobi",
	}
	fmt.Println(e.cost(),e.format())

	
	myEmail := email{isSubscribed: true, body: "Olá, tudo bem?"}
	printInfo(myEmail, myEmail) // passa como expense e formatter
}