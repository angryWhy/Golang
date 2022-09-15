package main

func main() {
	var arr [5]int = [5]int{1, 23, 4, 5} //数组必须指定长度和类型，且长度和类型指定后不可变
	var arr2 = [5]int{}
	var arr3 = [5]int{1, 23}
	var arr4 = [5]int{1: 2, 3: 0}   //指定index赋值
	var arr5 = [...]int{2, 3, 4, 5} //根据{}里的元素个数推断数组长度
	var arr6 = [...]struct {
		name string
		age  int
	}{{"tom", 18}, {"jim", 20}}
}

// 二维数组初始化
var arr = [5][3]int{{1, 2}, {3, 4, 5}} //五行三列
var arr3 = [...][3]int{{1, 2}, {3, 4}} //动态赋值行数

//访问元素
//index下标值访问
//首元素 arr[0]
//末元素 arr[len(arr)-1]

//cap和len

/*cap代表capacity容量
len代表length长度
len代表数组里有几个元素，cap代表分配给数组的内存了空间容纳多少个元素
有余数组初始化后长度不变，不需要预留空间，len(arr)==cap(arr)*/

//go没有按引用传参，是按值传参，传递的值实际上是数组的拷贝
//修改传进来的数组，不影响外部数组，如果修改，传入指针
