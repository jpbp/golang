/*
Unlike Python, Go is not function-scoped, it's block-scoped. Variables declared inside
a block are only accessible within that block (and its nested blocks).
There's also the package scope. We'll talk about packages later, but for now,
you can think of it as the outermost, nearly global scope.
*/

package main

import "fmt"

// scoped to the entire "main" package (basically global)
//var age = 19

func sendEmail() {
    // scoped to the "sendEmail" function
    name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
		fmt.Println(email)	
    }
	fmt.Println(name)
}

func splitEmail(email string) (string, string) {
	
	username, domain := "", ""
	
	for i, r := range email {
		if r == '.' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func main(){
	sendEmail()
	{
        age := 19
    //     // this is okay
        fmt.Println(age)
    }

    // this is not okay
    // the age variable is out of scope
    
	fmt.Println(splitEmail("rhaenyra@targaryen.com"))

}
