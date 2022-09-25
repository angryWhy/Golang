package main

import "fmt"

func main() {
	var sum = func(a, b int) int {
		return a + b
	}
	op(sum, 2)
	fmt.Println(op(sum, 2))
}

type add1 func(a, b int) int

func op(f add1, a int) int {
	b := a * 2
	return f(1, b)
}
