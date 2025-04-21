package main

import "fmt"

type User struct {
	Name string
	Membership
}

type Membership struct{
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	limit := 100
	if membershipType == "premium" {
		limit = 1000
	}
	
	user := User{
		 Name: name,
		 Membership: Membership {
			Type: membershipType,
			MessageCharLimit: limit,
		 },
	 }

	 return user
}

func main(){
	user1 := newUser("jp","teste")
	fmt.Println(user1)
}