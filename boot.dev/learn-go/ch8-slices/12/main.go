package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i,m :=range msg {
		for _,b :=range badWords{
			if (m==b){
				return i
			}
		}
	}
	return -1
}
func main(){
	b1 := [5]string{"crap","shoot","frick","dang"}
	m1 := [7] string{"what" , "the" , "shoot" , "I" , "hate" , "that" , "crap"}

	fmt.Println(indexOfFirstBadWord(m1[:],b1[:]))
}

	

 