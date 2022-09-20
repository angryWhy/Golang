package main

import "fmt"

func main() {
	color := "black"
	switch color {
	case "black":
		fmt.Println("ok")
		fallthrough
	default:
		fmt.Println("")
	}

}
