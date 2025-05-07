package main

import (
	"fmt"
	"time"
)

func main (){
	ch := make(chan string)

	go func(){
		time.Sleep(15 * time.Second)
		ch <- "Olá mundo"
	}()

	fmt.Println( <- ch)
}