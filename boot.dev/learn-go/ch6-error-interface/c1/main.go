package main

import (
	"errors"
	"fmt"
)
func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

func main(){
	s1 := "status exceeds 140 characters"
	s2 := ""
	s3 := "This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit."
	fmt.Println(validateStatus(s1))
	fmt.Println(validateStatus(s2))
	fmt.Println(validateStatus(s3))
}