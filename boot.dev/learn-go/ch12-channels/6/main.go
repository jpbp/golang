package main

import (
	"fmt"
)

func concurrentFib(n int) []int {
	fib := []int{}
	ch := make(chan int,n)
	
	go fibonacci(n,ch)
	for item := range ch{
		fmt.Println(item)
		fib = append(fib, item)
	}
	return fib

}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	fmt.Println("acabou")
	close(ch)
}
func main(){
	fmt.Println(concurrentFib(10))
}