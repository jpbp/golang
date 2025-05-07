package main

import (
	"fmt"
	"sync"
	"time"
)

func task (name string, wg *sync.WaitGroup){
	for i:=0;i<10;i++{
		fmt.Printf("%d: Task %s is running\n",i,name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}
//thread 1
func main(){
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	//thread 2
	go task("A",&waitGroup)
	//thread 3
	go task("B",&waitGroup)
	//thread 1
	go func ()  {
		for i:=0;i<10;i++{
			fmt.Printf("%d: Task Anonymous is running\n",i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}