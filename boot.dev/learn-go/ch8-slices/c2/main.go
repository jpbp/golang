package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	size := len(password)	
	if (size >= 5 && size <=12){
		upper:=0
		digit:=0
		for _,p := range password {
			if (unicode.IsUpper(p)){
				upper++
			}
			if(unicode.IsDigit(p)){
				digit++
			}
		}
		if (upper>0 && digit>0){
			return true
		}
	}
	return false
}

func main(){
	fmt.Println(isValidPassword("testW"))
}
