package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}


func main(){
	canal := make(chan int)
	go func (){
		canal <- 42
	}()
	fmt.Println(<-canal)


	e1 := email{body: "Words are pale shadows of forgotten names. As names have power, words have power.",date: time.Date(2010,12,1,0,0,0,0,time.UTC)}
	e2 := email{body: "Words are pale shadows of forgotten names. As names have power, words have power.",date: time.Now()}
	e3 := email{body: "Words are pale shadows of forgotten names. As names have power, words have power.",date: time.Now()}

	emails := []email{e1,e2,e3}
	//canal1 := make(chan bool)

	fmt.Println(checkEmailAge([3]email(emails)))

}