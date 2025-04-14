package main

import "fmt"

func main() {
	
	//s := []int{2,4,6,10}
	s := []int{10,20,30,40,50,60,70,80,90,100}

	fmt.Printf("len=%d cap=%d %v\n", len(s),cap(s),s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]),cap(s[:0]),s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]),cap(s[:4]),s[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]),cap(s[2:]),s[2:])

	s = append(s, 110)
	s = append(s, 120)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]),cap(s[2:]),s[2:])
	fmt.Printf("len=%d cap=%d %v\n", len(s),cap(s),s)
	
}
