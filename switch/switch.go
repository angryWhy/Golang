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
func switchType() {
	var num interface{} = 6.5
	switch num.(type) {
	case int:
		fmt.Printf("%d", num)
	case float64:
		fmt.Printf("%f", num)
	case bool:
		fmt.Printf("%b", num)
	}
}
