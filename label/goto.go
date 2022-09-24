package main

import "fmt"

func main() {
	gotoLabel()
	var i = 0
	if i%2 == 0 {
		goto L1
	}
	if i%2 == 1 {
		goto L2
	}
L1:
	i += 3
L2:
	i += 2
}
func gotoLabel() {
	s1 := make([]int, 3, 5)
	s2 := append(s1, 3)
	fmt.Printf("%v,%v", s1, s2)
}
