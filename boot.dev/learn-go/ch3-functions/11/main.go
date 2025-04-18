/*
Explicit Returns
Even though a function has named return values, we can still explicitly return values if we want to.

func getCoords() (x, y int){
  return x, y // this is explicit
}

Using this explicit pattern we can even overwrite the return values:

func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}

Otherwise, if we want to return the values defined in the function signature we can just use a naked return (blank return):

func getCoords() (x, y int){
  return // implicitly returns x and y
}

Assignment
Fix the bug in the code so that it returns the named values explicitly.
*/

package main

import "fmt"

func getCoords1() (x, y int)  {
	return x, y // this is explicit
}

func getCoords2() (x, y int) {
	return 5, 6 // this is explicit, x and y are NOT returned
}

func getCoords3() (x, y int) {
	return // implicitly returns x and y
}


func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return 
}

func main() {
	fmt.Println(getCoords1())
	fmt.Println(getCoords2())
	fmt.Println(getCoords3())
	fmt.Println(yearsUntilEvents(15))
}
