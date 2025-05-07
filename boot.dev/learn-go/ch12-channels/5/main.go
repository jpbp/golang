package main

import "fmt"

func countReports(numSentCh chan int) int {
	count := 0
	for {
		ch,ok := <- numSentCh
		if(!ok){
			break
		}
		count += ch
		fmt.Println(count,"countReports")
	}
	return count
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		fmt.Println(numReports,"sendReports")
		ch <- numReports
	}
	close(ch)
}
func main (){
	ch := make(chan int)
	go sendReports(10,ch)
	countReports(ch)
	//fmt.Println(ch,"main")
}