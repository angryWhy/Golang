package main

import (
	"fmt"
	"math/rand"
)

// 基本操作
func main() {
	var ch chan int
	fmt.Printf("%v\n", ch)
	ch = make(chan int, 10)
	fmt.Printf("%v,len=%d,cap=%d\n", ch, len(ch), cap(ch))
	ch <- 3
	ch <- 1
	fmt.Printf("%v,len=%d,cap=%d\n", ch, len(ch), cap(ch))
	<-ch
	fmt.Printf("%v,len=%d,cap=%d\n", ch, len(ch), cap(ch))
	makeSlice()
}

// 遍历
var ch = make(chan string, 5)

func forChan() {
	close(ch)
	ch <- "呵呵"
	ch <- "哈哈"
	for ele := range ch {
		fmt.Println(ele)
	}
}
func makeSlice() {
	s := makeS(100)
	fmt.Println(len(s))
	m := make(map[int]int, len(s))
	for _, v := range s {
		m[v] = 1
	}
	fmt.Println(len(m))
}
func makeS(n int) []int {
	//变换seed
	source := rand.NewSource(3333223)
	rander := rand.New(source)
	var s []int
	s = make([]int, 0, 10)
	for i := 0; i < n; i++ {
		s = append(s, rander.Intn(128))
	}
	return s
}
