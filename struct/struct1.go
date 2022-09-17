package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	time time.Time
	age  int
}

func (user User) say() int {
	return user.age
}
func (_ User) jump() {
	fmt.Println("hello")
}
func say(u User) {
	u.name = "zxc"
}
func say2(u *User) {
	u.name = "zxc"
}
func main() {
	u := User{
		name: "hello",
	}
	say(u)
	fmt.Printf("%s\n", u.name)
	say2(&u)
	fmt.Printf("%s\n", u.name)
}
