package main

import "fmt"

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

func getLast[Z any](s []Z) Z {
	var myZero Z
	if len(s) == 0{
		return myZero
	}
	return  s[len(s)-1]
}


func main(){
	array1 := []int{1,2,3,4,5,6,7,8,9,10}
	array2 := []string{"jp","ana","marco","mae","pitica"}

	fmt.Println(splitAnySlice(array1))
	fmt.Println(splitAnySlice(array2))
	fmt.Println(getLast(array1))
	fmt.Println(getLast(array2))
}

