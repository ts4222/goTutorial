package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
		if err == e1 {
			fmt.Println("err1")
		}
		if err == e2 {
			fmt.Println("err2")
		}
	}
}
