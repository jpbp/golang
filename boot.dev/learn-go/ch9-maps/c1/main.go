package main

import (
	"fmt"
	"strings"
)
func countDistinctWords(messages []string) int {
	countWords := make(map[string]int)
	for _,wordsRepeat := range messages {
		wordsRepeat = strings.ToLower(wordsRepeat)
		runes := []rune(wordsRepeat)
		var myString string
		for i, r := range runes{
			if (r != ' '){
				//fmt.Printf("%c",r)
				myString+=string(r)
				if (i == len(runes)-1){
					if _,ok := countWords[myString]; !ok{
						countWords[myString]=1
					}else{
						countWords[myString]++
					}	
				}
			}else{
				if _,ok := countWords[myString]; !ok{
					countWords[myString]=1
				}else{
					countWords[myString]++
				}
				myString=""
			}
		}
	}
	return len(countWords)
}

func main(){
	messages := []string{"WTS Arcanite Bar! Cheaper than AH","Do you need an Arcanite Bar!"}
	messages1 := []string{"LFG UBRS","lfg ubrs","LFG Ubrs"}
	fmt.Println(countDistinctWords(messages))
	fmt.Println(countDistinctWords(messages1))
	// wordsRepeat := "Could you give me a number crunch real quick?"
	// wordsRepeat = strings.ToLower(wordsRepeat)
	// runes := []rune(wordsRepeat)
	// var myString string
	// countWords := make(map[string]int)
	// for _, r := range runes{
	// 	if (r != ' '){
	// 		//fmt.Printf("%c",r)
	// 		myString+=string(r)
	// 	}else{
	// 		if _,ok := countWords[myString]; !ok{
	// 			countWords[myString]=1
	// 		}else{
	// 			countWords[myString]++
	// 		}
	// 		myString=""
	// 	}
		
	// }
	// fmt.Println(countWords)

}