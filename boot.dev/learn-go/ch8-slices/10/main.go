package main
import "fmt"
func sum(nums ...int) int {
	soma := 0
	for i:=0; i<len(nums); i++{
		soma += nums[i]
	}
	return soma
}

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}
func main(){
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)

	s := sum(1,2,3,4,5)
	fmt.Println(s)
}
