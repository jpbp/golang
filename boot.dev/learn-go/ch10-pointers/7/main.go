package main

import "fmt"

type circle struct{
	x int
	y int
	radius int
}

func (c *circle) grow(){
	c.radius *=2
}


func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}


func main (){
	c := circle{x: 1,y: 2,radius: 4}
	fmt.Println(c)
	c.grow()
	fmt.Println(c)


	newEmail := email{message: "teste123",fromAddress: "jp@teste.com",toAddress: "jp1@teste.com"}
	fmt.Println(newEmail)
	newEmail.setMessage("teste456")
	fmt.Println(newEmail)
}