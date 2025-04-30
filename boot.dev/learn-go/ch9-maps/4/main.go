package main
import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	
	countNames := make(map[rune]map[string]int)
	for _,name := range names{
		runes := []rune(name)
		if countNames[runes[0]] == nil{
			countNames[runes[0]]=make(map[string]int)
		}
		countNames[runes[0]][name]++
	}
	return countNames	
}

func main (){
	msg := []string{"billy","bino","billy","jp","ana","ana"}
	fmt.Println(getNameCounts(msg))
	
}
