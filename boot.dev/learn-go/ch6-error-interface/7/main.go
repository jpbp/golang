package main

import (
	"errors"
	"fmt"
)

type User struct{
	name string
	id string
}

func getUser(UserId string) (User,error){
	u := User {name: "jp",id: UserId}
	if UserId=="10"{
		return u, errors.New("Usuario nao existe")
	}
	return u,nil
}

func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}