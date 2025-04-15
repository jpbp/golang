/*

The Initial Statement of an If Block
An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

if INITIAL_STATEMENT; CONDITION {
}

Why Would I Use This?
It has two valuable purposes:

It's a bit shorter
It limits the scope of the initialized variable(s) to the if block
For example, instead of writing:

length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}

We can do:

if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}

In the example above, length isn't available in the parent scope, which is nice because we don't need it there - we won't accidentally use it elsewhere in the function.
*/

package main

import "fmt"

func main() {
	email := "Bom dia Prezado!!! Estou mandando este email para ajudar no nosso estudo sobre Go. Espero que esteja tudo bem com você. Um abraço."
	//var email string
	length := getLength(email)
	if length < 1 {
		fmt.Println("Email is invalid")
	}else{
		fmt.Println("Email is valid")
	}
	
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	}else{
		fmt.Println("Email is valid")
	}

}

func getLength(email string) int {
	return len(email)
}
