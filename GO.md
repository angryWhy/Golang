## Golang

## 数据类型

### 基础数据类型

```go
bool		 1字节		默认值：false
byte		 1字节		默认值：uint8，取值范围[0,255]
rune		 4字节		unicode ponit code int32
int，uint     4或8字节   取决于操作系统，32或64
int8,uint8	  1字节	 -128~127，0~255
int16，uint16 2字节	
float32
float64
complex64				  对应的实部或者虚部：float64
complex128				  对应的实部或者虚部：float128
uintptr
```

### 复合数据类型

```go
array		默认值：无		值类型
struct		默认值：无		值类型
string		默认值：""		 utf-8字符串
slice		默认值：nil		 引用类型
map			默认值：nil		 引用类型
channel		默认值：nil		 引用类型
interface	默认值：nil		  接口
function	默认值：nil		  函数
```

### 指针

```go
指针方式
a := 123
var poniter unsafe.Pointer = unsafe.Pointer(&a)
var p  uintptr = uintptr(poniter)
var ptr *int = &a
//取址
a :=1
var b = &a
//解析指针
*b = 2
```

### 自定义类型

#### 类型别名

```go
type myint = int8
type myfloat = float32
```

#### 自定义类型

```go
type ms map[string]string
type add func(a,b int)int
type user struct {name:int;age:int}
```

#### 结构体

```go
type User struct{
    Name string
    Age int
}
func (ms User) hello(){
    fmt.Printf("my name is %s/n",ms.Name)
}
//赋值
var u User
u = User{
	"tom",
     18
}
u.hello()
```

### 字符串

```go
s :="my name " //字符串可以包含Unicode字符
s1 :="my name \\" //包含转义字符
s2 :=`name    hello`//原封不动输出，包括空白符合换行符
```

###### 格式化输出

```go
 		   %v,原样输出
            %T，打印类型
            %t,bool类型
            %s，字符串
            %f，浮点
            %d，10进制的整数
            %b，2进制的整数
            %o，8进制
            %x，%X，16进制
                %x：0-9，a-f
                %X：0-9，A-F
            %c，打印字符
            %p，打印地址
```



#### 常用API

```go
len(str)							求长度
strings.Split						分割
strings.Contains					判断是否包含
strings.HasPrefix,strings,HasSuffix   前缀后缀判断
strings.Index(),strings.LastIndex     寻找子字符串出现的位置
//拼接
func fmt.Sprintf(format string ,a ...interface{}) string
//参数一：返回拼接后的字符串
func strings.Join(elems []string,sep string) string
//参数二：字符串数组，中间的连接符
//当有大量字符串进行拼接的时候，使用builders
```

```go
string每个元素叫做字符，字符有两种
	byte：1字节，代表ASCILL码的一个字符
	rune：4个字节，一个代表一个UTF-8字符，一个汉字代表一个rune
string底层是byte数组，string的长度就是byte数组的长度
string可以转换成byte[]，或者rune[]类型
string是常量，不可被修改
//
func main() {
	s := "hahaha 123 www"
	arr := []byte(s)
	for _, ele := range arr {
		fmt.Printf("%d", ele)
	}
	fmt.Print("arr len %d,s len %d\n", len(arr), len(s))
}
//
func main() {
	s := "1一"
    //字符串长度
    one :=len([]rune(s))
    //[49,228,184,128]
    two :=len([]byte(s))
}
```

###### split

```go
func str_spl(){
   	s := "hello go"
	arr := strings.Split(s, "o")
	fmt.Println(arr)
}
```

###### Contains

```go
func str_spl(){
   	s := "hello go"
	arr := strings.Contains(s, "o")
	fmt.Println(arr)
}
```

###### HasPrefix,HasSuffix

```go
func str_spl() {
	s := "hello go"
	arr := strings.HasPrefix(s, "e")
	fmt.Println(arr)
}
func str_spl() {
	s := "hello go"
	arr := strings.HasSuffix(s, "o")
	fmt.Println(arr)
}
```

###### Index

```go
func str_spl() {
	s := "hello go"
	arr := strings.Index(s, "o")
	fmt.Println(arr)
}
```

###### 字符串拼接

```go
func add() {
	s := "hello"
	t := "ni"
	p := "hao"
	sub1 := s + "" + t + "" + p
	sub2 := fmt.Sprintf("%s %s %s", s, t, p)
	sub3 := strings.Join([]string{s, t, p}, "")
	sub4 := strings.Builder{}
	sub4.WriteString(s)
	sub4.WriteString("")
	sub4.WriteString(t)
	sub4.WriteString("")
	sub4.WriteString(p)
	sub5 := sub4.String()
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
	fmt.Println(sub5)

}
```

###### 强制类型转换

```go
byte和int可以相互转换
float和int可以转换，小数位会丢失
bool和int不能相互转换
不同长度的int或float转换
string可以转换成[]byte或[]rune类型
低精度向高精度转换可以，高向低转换丢失位数
无符号向有符号转换，最高位是符号位
```

###### 丢失例子

```go
func test() {
	var ua uint64 = math.MaxUint64
	fmt.Printf("ua=%d %b\n", ua, ua)
	ue := uint8(ua)
	fmt.Printf("ue=%d %b\n", ue, ue)
	ub := uint64((ue))
	fmt.Printf("ub=%d %b\n", ub, ub)
}
//ua=18446744073709551615 1111111111111111111111111111111111111111111111111111111111111111
//ue=255 11111111
//ub=255 11111111
//发生丢失
```

### 数组

##### 初始化

```go
var arr1 [5]int = [5]int{}//数组必须指定长度和类型，且长度和类型指定后不可变
var arr2 = [5]int{}
var arr3 = [5]int{3,2}
var arr4 = [5]int{2:15,4:30}//指定index赋值
var arr5 = [...]int{3,2,6,4}//根据{}里的元素个数推断数组长度
var arr6 = [...]struct {
	name string
    age int
}{{"tom",18},{"jim",20}}
```

##### 二维数组初始化

```go
var arr = [5][3]{{1,2},{3,4,5}}//五行三列
var arr2 = [...][1]{{3,3,3},{1,1,1},{2,2,}}//动态赋值行数
```

##### 访问元素

```go
通过index访问
arr[0]//首元素
arr[len(arr)-1]

//访问二维数组
arr[1][2]//第二行第三列
```

##### 遍历数组

```go
var arr = [5]int{1,2,3,4,5}
//range
for i,ele := range arr{
    fmt.Println(ele)
}
//for
for i :=0;i<10;i++{
	fmt.Println(i)
}
//遍历二维数组
for i,ele :=range arr {
    for e,i := range ele{
    	fmt.Println(arr[i][e])
    }
}
```

##### cap和len

```go
cap代表capacity容量
len代表length长度
len代表数组里有几个元素，cap代表分配给数组的内存了空间容纳多少个元素
有余数组初始化后长度不变，不需要预留空间，len(arr)==cap(arr)
```

#####  数组传参

```
go没有按引用传参，是按值传参，传递的值实际上是数组的拷贝
修改传进来的数组，不影响外部数组，如果修改，传入指针
```

### 切片

```go
//切片的地址和地址首元素的地址是两回事
type slice struct {
	array unsafe.Pointer //指针
    len int				//长度
    cap int				//容量
}
```

##### 初始化

```go
var s []int			//切片声明，len=cap=0
s = []int{}			//初始化，len=cap=0
s = make([]int,3)	//初始化，len=cap=3
s = make([]int,3,5)	//初始化，len=3,cap=5
s = []int{1,2,3,4,5}//初始化，len=cap=5
s2d :=[][]int{
    {1},{2,3}
}
```

##### append

```go
切片相对于数组，可以自动扩容，可以追加元素
追加元素放在预留空间里，长度加1
预留空间用完，则会申请新的内存空间，新申请的内存空间并且扩容（内存对齐），把原内存空间拷贝出来，执行append操作
```

```go
func main() {
	var s []int
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{}
	s = make([]int, 3, 5)
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = append(s, 999)
	fmt.Printf("len %d cap %d %d\n", len(s), cap(s), s)
}
```

##### 子切片

```go
//长3，容量5
s :=make([]int,3,5)
//截取1,2位元素,len2,cap4(首地址是1，对于母切片上，有1,2,3,4容量)
sub_slice = s[1:3] 
//子切片和母切片共享底层的内存空间，子切片修改会反映到母切片上，在子切片上执行append会把新元素放到母切片的预留内存空间上
//当子切片不断地append，消耗完母切片预留的内存空间，子切片和母切片发生内存分离，两个切片毫无关系
```

#### 练习

```go
实现一个func arrstring(arr []int) string(){
	//比如输入[]int{1,2,3}，返回“1,2,3”,切片很短，也可能很长
}
func addStr(n []int) string {
	var s strings.Builder
	for _, ele := range n {
		s.WriteString(strconv.Itoa(ele))
		s.WriteString(" ")
	}
	return strings.Trim(s.String(), " ")
}
```

### map

```go
//go map底层实现是hash table，根据key查找value的时间复杂度是O(1)
//key值生成hash值，对数组长度取模，然后分配到对应槽位，如有冲突，对应的槽位上形成链表
//如果槽位利用率过高，则会重新分配，key
```

初始化

```go
var m map[string]int			//声明map，指定的key和value的数据类型
m = make(map[string]int)		//初始化，容量为0
m = make(map[string]int,5)		//初始化，容量为5，最好给个合适的容量
m = map[string]int{"haha":2,"hehe":3}
m["h"] = 1
m["h"] = 2//会覆盖之前的值
delete(m,"h")//删除元素
len(m)//获取map的长度
//不能调用cap函数

//读取值
value :=m["h"]
//如果不存在，默认值，例如int，则会是0
要判断
value,ok :=m["O"]
```

```go
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
```

#### 练习

```go
创建一个长度为100的切片，添加100个元素，map去重
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
	var s []int
	s = make([]int, 0, 10)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(128))
	}
	return s
}
```

### channel

管道是一个环形队列（先进先出），send（插入），recv（取走），从同一个位置，向同一个方向顺序执行

sendx代表最后插入元素的位置，recvx代表最后一次取走元素的位置

##### 初始化

```go
var ch chan int				//声明
ch = make(chan int,8)		//初始化，只能make函数，环形队列容量8个，不能扩容
```

##### 存放数据

```go
ch <-1 //向管道里写数据
ch <-5
ch <-6
v :=<-ch//向管道里取数据
s :=<-ch
```

##### 遍历管道

```go
close(ch)			//遍历前，先关闭管道，禁止写入数据
for ele := range ch {
		fmt.Println(ele)
	}

//for循环
先保存长度
length :=len(ch)
for i:=0;i<length;i++{}
//循环完长度为0
相当于取出来进行操作
```

### 结构体

##### 定义

```go
type user struct{
	id int
	score float32
	time time.Time
	name,address string
}
```

##### 初始化一个实例

```go
var u user //相对应的默认值初始化struct
u = user{}
u = user{id:1,score:2.2,}//赋值
u = user{4,100.00}//必须按照顺序赋值
//访问
u.id
```

##### 成员函数（方法）

```go
func(u User) hello(t string){
	fmt.println(u.name)
}
//不使用
func (_ User)hello(t string){
	fmt.println(u.name)
}
```

##### 任意类型添加方法

```go
type UserMap map[string]User
func (um UserMap)GetUser(id int)int{
	return um[id]
}
```

##### 可见性

```
字母大写，跨package访问
```

##### 匿名结构体

只是用一次的情况

```
var user struct{
	id int
}
```

##### 例子

```go
type User struct {
	name string
	time time.Time
	age  int
}

func (user User) say() int {
	return user.age
}
func (_ User) jump() {
	fmt.Println("hello")
}
func main() {
	u := User{
		name: "wzx",
		time: time.Now(),
		age:  22,
	}
}
```

##### 结构体指针

```go
type User struct {
	name string
	time time.Time
	age  int
}
var u User
//1.
user := &u
//2.如果赋值了，就不等价于1.
user = &User{name:"he",......}
//3.通过new函数实体化一个结构体，返回其指针
user = new(User)
```

构造函数指针

好处，返回地址，不用拷贝值

```
func NewUser() *User{
	return &User{
		id:1,
		name:@
	}
}
```

```go
//user传的是值，即是整个结构体的拷贝，在函数里修改结构体，不会影响原来的结构体
func hello (u User,man string){
	u.name = "hello jack"
}
//传的是User的指针，可以原先的结构体
func say (u *User,man string){
	u.name = "hello jack"
}

//
func say(u User) {
	u.name = "zxc"
}
func say2(u *User) {
	u.name = "zxc"
}
func main() {
	u := User{
		name: "hello",
	}
	say(u)
	fmt.Printf("%s\n", u.name)
	say2(&u)
	fmt.Printf("%s\n", u.name)
}
```

#### 结构体嵌套

嵌套：逐级访问属性

对于匿名结构体可以跳过，直接访问内部属性

如果冲突,加上中间字段名字

```go
type user struct{
	name string
}
type person struct{
	student user
}
//匿名字段
type teacher struct{
    name int
    //匿名
    user
}
//冲突，例如
t :=new (teacher)
t.name = "hello"
```

### 深浅拷贝

深拷贝：开辟新的内存空间

浅拷贝：拷贝的是指针

```go
slice引用类型，底层实现了：数组的指针，len，cap
user :=[]user{{name:"haha"}}
func main(user []){
	user[0].name = "hehe"
}
拷贝的底层数组：修改会反映到原来的slice上
```

### if语句

```go
if语句
if 9>5 {
	//执行这里的内容
}
//可以包含初始化变量,表达式可以包含变量或者常量，允许包含一个分号，变量在if块可见
if c,d :=1,3;c<d{
	
}
//例如使用map
	m := make(map[string]int, 10)
	if value, exist := m["1"]; exist {
		fmt.Println(value)
	}
```

### switch语句

```go
	color := "black"
	switch color {
		case "black":
			fmt.Println("ok")
       		 fallthrough
		default:
			fmt.Println("")
	}
//fallthrough
case里面带了fallthrough，会继续执行后面的代码
```

