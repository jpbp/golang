package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user,ok := users[name]
	if !ok {
		return false,errors.New("not found")
	}

    if user.scheduledForDeletion == false{
		return false,nil
	}

	delete(users,name)
	return true,nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}


func main(){
	users := make(map[string]user)

	users["jp"]= user{name: "jp",number: 122345,scheduledForDeletion: false}
	users["ana"]= user{name: "ana",number: 456123,scheduledForDeletion: true}
	users["marco"]= user{name: "marco",number: 789123,scheduledForDeletion: false}
	for _,u := range users{
		fmt.Println(u)
	}

	fmt.Println(deleteIfNecessary(users,"jp"))
	fmt.Println(deleteIfNecessary(users,"ana"))
	fmt.Println(deleteIfNecessary(users,"marco"))

	for _,u := range users{
		fmt.Println(u)
	}
}
