/*
Conditionals
if statements in Go do not use parentheses around the condition:

if height > 4 {
    fmt.Println("You are tall enough!")
}

else if and else are supported as you might expect:

if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}

Assignment
Fix the bug on line 12. If messageLen is less than or equal to the maxMessageLen the program should print "Message sent", else it should print "Message not sent".

Tips
Here are some of the comparison operators in Go:

== equal to
!= not equal to
< less than
> greater than
<= less than or equal to
>= greater than or equal to
*/
package main
import "fmt"
func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}