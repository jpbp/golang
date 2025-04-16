/*
Passing Variables by Value
Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.

func main() {
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int) {
    x++
}

Assignment
Fix the bugs in the monthlyBillIncrease and getBillForMonth functions.

monthlyBillIncrease: Returns the increase in the bill from the previous to the current month. If the bill decreased, it should return a negative number.
getBillForMonth: Returns the bill for the given month.
It looks like whoever wrote the getBillForMonth function thought that they could pass in the bill parameter, update it inside the function, and that update would apply in the parent function (monthlyBillIncrease). They were wrong.

Change the getBillForMonth function to explicitly return the bill for the given month, and be sure to capture that return value properly in the monthlyBillIncrease function.

The function signature for getBillForMonth should only take 2 parameters once you're done.
*/
package main
import "fmt"
func main() {
    x := 5
    increment(&x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x

	fmt.Println(monthlyBillIncrease(15,20,50))
}

func increment(x *int) {
    *x++
}


func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
}
