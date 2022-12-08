package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func GetUser() (user *User) {
	return
}

func main() {
	user := GetUser()
	if user == nil {
		fmt.Println("nil panic")
	}
	fmt.Println(user)
}
