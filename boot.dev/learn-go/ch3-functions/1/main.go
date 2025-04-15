/*
Functions
Functions in Go can take zero or more arguments.

To make code easier to read, the variable type comes after the variable name.

For example, the following function:

func sub(x int, y int) int {
  return x-y
}

Accepts two integer parameters and returns another integer.

Here, func sub(x int, y int) int is known as the "function signature".

Assignment
We often will need to manipulate strings in our messaging app. For example, adding some personalization by using a customer's name within a template. The concat function should take two strings and smash them together.

hello + world = helloworld
Fix the function signature of concat to reflect its behavior.
*/

package main

import "fmt"

func sub (x int ,  y int ) int{
	return x-y
}

func concat (s1 string, s2 string) string{
	return s1 + s2
}

func main (){
	test("Lane,", " happy birthday!")
	test("Zuck,", " hope that Metaverse thing works out")
	test("Go", " is fantastic")
	test2(10,5)
	test2(12,6)
	test2(16,5)
}

func test (s1 string, s2 string) {
	fmt.Println(concat(s1,s2))
}

func test2 (x int, y int){
	fmt.Println(sub(x,y))
}