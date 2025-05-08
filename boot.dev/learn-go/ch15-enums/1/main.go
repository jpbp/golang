package main

import "fmt"

type perm string

const (
    Read  perm = "read"
    Write perm = "write"
    Exec  perm = "execute"
)

var Admin = "admin"
var User = perm("user")

func checkPermission(p perm) {
    fmt.Println(p)
}

func main (){
	checkPermission(perm(Admin))
}