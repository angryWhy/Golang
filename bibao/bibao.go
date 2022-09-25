package main

import "fmt"

func main() {
	f := add()
	f()
	f()
}
func add() func() {
	i := 2
	var b = func() {
		fmt.Printf("i的指针%p\n", &i)
		i--
		fmt.Printf("i的值%d\n", i)
	}
	return b
}
