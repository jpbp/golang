package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	mutex := &sync.Mutex{}
	go writeLoop(m,mutex)
	go readLoop(m,mutex)
	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mutex *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}
	}
}

func readLoop(m map[int]int,mutex *sync.Mutex) {
	for {
		mutex.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mutex.Unlock()
		
	}
}