package main

import "fmt"

//thread1
func main(){
	canal := make(chan string) //vazio
	//thread2
	go func ()  {
		canal <- "Olá Mundo!" // cheio
		canal <- "Olá Mundo2!" // cheio
	}()

	msg:= <-canal // canal vazio
	fmt.Println(msg)
}