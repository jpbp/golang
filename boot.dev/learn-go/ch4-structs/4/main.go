package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func main(){
	sender1 := sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 12348484,
		},
	}
	fmt.Println(sender1.name)   // "Deborah"
	fmt.Println(sender1.number) // 12348484
	fmt.Println(sender1.rateLimit)
}