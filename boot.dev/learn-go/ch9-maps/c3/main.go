package main

import "fmt"

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]
	defer delete(users, name)
	if !ok {	
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

func main(){
	users := make(map[string]user)
	u1 := user{name: "jp",number: 12345678,admin: true}
	u2 := user{name: "ana",number: 12345678,admin: false}
	u3 := user{name: "marco",number: 12345678,admin: true}
	users["jp"]=u1
	users["ana"]=u2
	users["marco"]=u3
	
	fmt.Println(logAndDelete(users,"jp"))
	fmt.Println(logAndDelete(users,"ana"))
	fmt.Println(logAndDelete(users,"marc"))
}
