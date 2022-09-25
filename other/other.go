package main

import "fmt"

func main() {
	brr := []byte{1, 2, 3}
	update(brr...)
	update(4, 5, 3)
}
func update(other ...byte) {
	for _, ele := range other {
		fmt.Println(ele)
	}
}
