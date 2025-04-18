/*

Functions As Values
Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.

Let's assume we have two simple functions:

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

We can create a new aggregate function that accepts a function as its 4th argument:

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}

It calls the given arithmetic function (which could be add or mul, or any other function that accepts two ints and returns an int) and applies it to three inputs instead of two. It can be used like this:

func main() {
	sum := aggregate(2, 3, 4, add)
	// sum is 9
	product := aggregate(2, 3, 4, mul)
	// product is 24
}


*/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}



func addExclamation(text string) string{
	return text+"!"
}
func addPeriod(text string) string{
	return text+"."
}
func reverseString(s string) string { 
    rns := []rune(s) // convert to rune 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        // swap the letters of the string, 
        // like first with last and so on. 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    // return the reversed string. 
    return string(rns) 
} 
func reformat(message string, formatter func(string) string) string {
	for i := 0; i<3; i++{
		message = formatter(message)
	}
	return "Textio: "+message
}

func main() {
	sum := aggregate(2, 3, 4, add)
	// sum is 9
	product := aggregate(2, 3, 4, mul)
	// product is 24

	fmt.Println(sum,product)
	fmt.Println(addExclamation("hello"))
	fmt.Println(addPeriod("hello"))
	fmt.Println(reverseString("hello"))
	fmt.Println(reformat("hello", addExclamation))
	fmt.Println(reformat("hello", addPeriod))
	fmt.Println(reformat("hello", reverseString))
}
