/*
Named Return Values
Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

According to the tour of go:

A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

Use naked returns if it's obvious what the purpose of the returned values is. Otherwise, use named returns for clarity.

func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}

Is the same as:

func getCoords() (int, int){
  var x int
  var y int
  return x, y
}

In the first example, x and y are the return values. At the end of the function, we could simply write return to return the values of those two variables, rather than writing return x,y.

Assignment
One of our clients likes us to send text messages reminding users of life events coming up.

Fix the bug by using named return values in the function signature. The variables need to be automatically initialized. Order them as they appear in the code. Do not alter the body of the function.
*/
package main
import "fmt"

func getCoords1() (x, y int){
	// x and y are initialized with zero values
  
	return // automatically returns x and y
}

func getCoords2() (int, int){
	var x int
	var y int
	return x, y
}
  

func yearsUntilEvents1(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {

	
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

func yearsUntilEvents2(age int) (int,int, int) {
	var yearsUntilAdult int
	var yearsUntilDrinking int
	var yearsUntilCarRental int
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
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main (){
	x1,x2 := getCoords1()
	x3,x4 := getCoords2()
	fmt.Println(x1,x2)
	fmt.Println(x3,x4)
	fmt.Println(yearsUntilEvents1(25))
	fmt.Println(yearsUntilEvents2(25))
}