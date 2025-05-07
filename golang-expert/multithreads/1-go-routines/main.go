package main

import (
	"fmt"
	"time"
)

func task (name string){
	for i:=0;i<10;i++{
		fmt.Printf("%d: Task %s is running\n",i,name)
		time.Sleep(1 * time.Second)
	}
}
//thread 1
func main(){
	//thread 2
	go task("A")
	//thread 3
	go task("B")
	//thread 1
	go func ()  {
		for i:=0;i<10;i++{
			fmt.Printf("%d: Task Anonymous is running\n",i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}