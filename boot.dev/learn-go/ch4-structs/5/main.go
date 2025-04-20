package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}


func (auth authenticationInfo) getBasicAuth() string{
	return "Authorization: Basic "+ auth.username+":"+auth.password
}

func getBasicAuth(auth authenticationInfo) string{
	return "Authorization: Auth+Pass "+ auth.username+":"+auth.password
}

func main(){
	auth := authenticationInfo{
		username: "Bing",
		password: "987456",
	}
	fmt.Println(auth.getBasicAuth())
	fmt.Println(getBasicAuth(auth))
}