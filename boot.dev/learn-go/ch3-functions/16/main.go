/*
Anonymous Functions
Anonymous functions are true to form in that they have no name. They're useful when defining a function that will only be used once or to create a quick closure.

Let's say we have a function conversions that accepts another function, converter as input:

	func conversions(converter func(int) int, x, y, z int) (int, int, int) {
		convertedX := converter(x)
		convertedY := converter(y)
		convertedZ := converter(z)
		return convertedX, convertedY, convertedZ
	}

We could define a function normally and then pass it in by name... but it's usually easier to just define it anonymously:

	func double(a int) int {
	    return a + a
	}

	func main() {
	    // using a named function
		newX, newY, newZ := conversions(double, 1, 2, 3)
		// newX is 2, newY is 4, newZ is 6

	    // using an anonymous function
		newX, newY, newZ = conversions(func(a int) int {
		    return a + a
		}, 1, 2, 3)
		// newX is 2, newY is 4, newZ is 6
	}

Assignment
Complete the printReports function. It takes as input a sequence of messages, intro, body, outro. It should call printCostReport once for each message.

For each call of printCostReport, give it an anonymous function that returns the cost of a message as an integer. Here are the costs:

Intro: 2x the message length
Body: 3x the message length
Outro: 4x the message length
Use the built-in len() function to get the length of a string:

helloLen := len("hello")
// helloLen = 5
*/
package main

import "fmt"

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func double(a int) int {
	return a + a
}

func printReports(intro, body, outro string) {
	printCostReport( func ( s string ) int {
		return 2 * len(s)
	} , intro)

	printCostReport( func ( s string ) int {
		return 3 * len(s)
	} , body)

	printCostReport( func ( s string ) int {
		return 4 * len(s)
	} , outro)
	
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func main() {
	// using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6
	fmt.Println(newX,newY,newZ)
	newX, newY, newZ = conversions(func(a int) int {
	    return a + a
	}, 1, 2, 3)
	fmt.Println(newX,newY,newZ)
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)

}
