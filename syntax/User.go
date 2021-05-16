package main

import "fmt"

type User struct {
	username string
	password string
}

func (p *User) init(username string, password string) (*User, string) {
	if "" == username || "" == password {
		return p, p.Error()
	}
	p.username = username
	p.password = password
	return p, ""
}

func (p *User) Error() string {
	return "Usernam or password shouldn't be empty!"
}

func main() {
	var user User
	user1, _ := user.init("", "")
	fmt.Println(user1)
}
