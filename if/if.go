package main

import "fmt"

func main() {
	if c, d := 1, 2; c < d {
		fmt.Println("ok")
	}
	m := make(map[string]int, 10)
	if value, exist := m["1"]; exist {
		fmt.Println(value)
	}
}
