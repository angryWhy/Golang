package main

import (
	"fmt"
	"time"
)

type PathError struct {
	name    string
	path    string
	time    string
	message string
}

func NewPathError(path, name string) PathError {
	return PathError{
		name:    name,
		path:    path,
		time:    time.Now().Format("2022-09-25"),
		message: "haha",
	}
}
func (e PathError) Error() string {
	return e.time + e.name + e.message
}
func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了panic，不让程序退出")
		}
	}()
	panic(2)
}
func main() {

}
