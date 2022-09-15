package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int, 10) //cap=10
	m = map[string]int{"a": 10, "b": 30}
	m["D"] = 10
	delete(m, "D")
	for key, value := range m {
		fmt.Printf("key=%s,value=%d\n", key, value)
	}
	value, ok := m["A"]
	if ok {
		fmt.Printf("%d\n", value)
	} else {
		fmt.Println("不存在")
	}
}

//go map底层实现是hash table，根据key查找value的时间复杂度是O(1)
//key值生成hash值，对数组长度取模，然后分配到对应槽位，如有冲突，对应的槽位上形成链表
//如果槽位利用率过高，则会重新分配，key
/*
m["h"] = 1
m["h"] = 2//会覆盖之前的值
delete(m,"h")//删除元素
len(m)//获取map的长度
//不能调用cap函数

//读取值
value :=m["h"]
*/
