package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	users := make(map[string]user)
	if (len(names) != len(phoneNumbers) ){
		return users,errors.New("invalid sizes")
	}
	for i,name := range names{
		users[name] = user {
			name: name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return users,nil
}

type user struct {
	name        string
	phoneNumber int
}

func main(){
	names := []string{"jp","ana","marco"}
	phoneNumbers := []int{12222,45555,78888}

	fmt.Println(getUserMap(names,phoneNumbers))

}