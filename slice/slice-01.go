package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 2, 3, 3, 4}
	arr2 := arr[0:2]
	fmt.Printf("母切片%p，子切片%p\n", &arr, &arr2)
	sub()
	append2()
}

// 切片的初始化
func init2() {
	var s []int
	s = []int{}
	s = make([]int, 5)     //指定长度
	s = make([]int, 5, 10) //len =5,cap =10
	s = []int{1, 23, 34}   //初始化
	fmt.Println(s)
}

// 子切片地址和切片首元素地址不是一样的
func sub() {
	s := make([]int, 6)
	sub := s[1:3]
	fmt.Printf("母切片的地址%p,首元素地址%p,子切片的地址%p,首元素地址%p", &s, &s[0], &sub, &sub[0])
}

// 切片相对于数组，可以自动扩容，可以追加元素
// 追加元素放在预留空间里，长度加1
// 预留空间用完，则会申请新的内存空间，新申请的内存空间并且扩容（内存对齐），把原内存空间拷贝出来，执行append操作
func append2() {
	s := make([]int, 5, 5)
	sub := s[1:3]
	fmt.Printf("母切片的长度%d,容量%d---子切片的长度%d，容量%d\n", len(s), cap(s), len(sub), cap(sub))
	sub = append(sub, 9)
	sub = append(sub, 9)
	sub = append(sub, 9)
	//子切片和母切片共享底层的内存空间，子切片修改会反映到母切片上，在子切片上执行append会把新元素放到母切片的预留内存空间上
	//当子切片不断地append，消耗完母切片预留的内存空间，子切片和母切片发生内存分离，两个切片毫无关系
	fmt.Printf("母切片的长度%d,容量%d---子切片的长度%d，容量%d\n", len(s), cap(s), len(sub), cap(sub))
}
