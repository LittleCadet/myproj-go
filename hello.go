package main // 这是必须的，声明包名

import (
	"errors"
	"fmt"
	"math/cmplx"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var age string

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// var可以声明所有变量
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(target_url)
	fmt.Println("go, Hello, World!")

	// var 与指定的数据类型配合使用
	var int8Num int8 = 123
	// var 可以一次性声明多个变量
	var a, b int = 8, 9
	var c, d = 8, "9"
	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(d)
	age = "我是age"
	// 类型转换挺麻烦： 没事别转： 以下是错误的
	// s := strconv.FormatInt(int64(int8Num), 10) // 10 表示十进制
	// fmt.Println(s)
	fmt.Println(int8Num)

	// 变量声明的另一种方式：
	fmt.Println("============多个变量同时赋值的多种方式：")

	// 相当于：
	// var intVal int =1
	intVal := 10
	a1, b1, c1 := 1, 2, "3"
	fmt.Println(intVal)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	// 值拷贝
	fmt.Println("==============值拷贝")
	var code1 = 1
	var code2 = code1
	fmt.Println(code1)
	fmt.Println(code2)

	// code1 = 2
	// fmt.Println(code1)
	// fmt.Println(code2)

	// 引用拷贝
	fmt.Println("==============引用拷贝")
	code3 := 123
	code4 := code3
	fmt.Println(code3)
	fmt.Println(code4)

	fmt.Println("=============复数")
	// 创建复数
	aa := 1 + 2i
	bb := complex(3.0, -1.0) // 等价于 3 - 1i

	// 复数运算
	sum := aa + bb
	product := aa * bb

	fmt.Println("实部：", real(aa), ", 虚部：", imag(aa))
	fmt.Println("aa =", aa)           // (1+2i)
	fmt.Println("bb =", bb)           // (3-1i)
	fmt.Println("aa + bb =", sum)     // (4+1i)
	fmt.Println("aa * bb =", product) // (5+5i)

	// 使用 math/cmplx 包
	magnitude := cmplx.Abs(aa)  // 模长 |a| = sqrt(1² + 2²)
	conjugate := cmplx.Conj(aa) // 共轭复数 (1-2i)

	fmt.Printf("|aa| = %.2f\n", magnitude) // 2.24
	fmt.Println("共轭 aa* =", conjugate)     // (1-2i)

	fmt.Println("===============常量")

	const (
		ac     = "这是一个常量abc"
		bc, dc = len(ac), unsafe.Sizeof(ac)
	)

	fmt.Println("ac:", ac, ", bc:", bc, ", dc:", dc)

	// iota ：是在const中的行索引， 第一个是0 ， 后面 +1
	const (
		// a10 = 100
		a11 = iota
		// a12 = iota
		// a13 = iota
		a12
		a13
	)

	fmt.Println("iota:", a11, a12, a13)

	fmt.Println("==============循环")

	count := 1
	for true {
		count++
		if count > 10 && count < 50 {
			continue
		}

		fmt.Println("这是循环:", count)

		if count > 100 {
			break
		}

		if count == 60 {
			goto Found
		}

	}

Found:
	fmt.Println("跳出循环")

	fmt.Println("==========函数调用")

	// var num1 = 1
	// var num2 = 2
	// var result int
	// result = sum(num1, num2)
	num1, num2 := 1, 2
	result := sum2(num1, num2)
	fmt.Println("调用 sum 函数：", result)

	// var a100 int = 100
	// var b100 int = 200
	a100, b100 := 100, 200

	/* 调用函数并返回最大值 */
	var ret100 = max(a100, b100)
	fmt.Println("调用 max 函数：", ret100)

	result2, result3 := swap(num1, num2)
	fmt.Println("返回多个参数：", result2, result3)

	fmt.Println("============数组")
	var dataArray1 [10]int
	dataArray2 := [10]int{1, 2, 3, 4, 5}
	var dataArray3 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataArray4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i, j int
	for i = 0; i < 10; i++ {
		dataArray1[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Println("dataArray1, dataArray2, dataArray3, dataArray4:", dataArray1[j], dataArray2[j], dataArray3[j], dataArray4[j])
	}

	fmt.Println("==========指针")

	var pointer = 1
	fmt.Println("pointer的指针：", &pointer)

	var pointer2 *int
	pointer2 = &pointer
	fmt.Println("pointer2的指针：", pointer2)

	fmt.Println("pointer2的值:", *pointer2)

	pointer = 2
	fmt.Printf("pointer,pointer2:", pointer, *pointer2)

	fmt.Println("===============实体类：struct:")

	// 注意： 用 {} 而不是 （）
	fmt.Println("book1:", Books{"书名", "作者", "科目", 1})
	var book2 Books = Books{"书名2", "作者", "科目", 2}
	fmt.Println("book2:", book2)

	var book4 = Books{title: "书名4", subject: "科目", author: "作者", book_id: 4}
	var book5 = Books{subject: "科目"}
	fmt.Println("book4, book5:", book4, book5)

	var book3 Books
	book3.author = "作者"
	book3.book_id = 3
	book3.subject = "科目"
	book3.title = "书名3"
	fmt.Println("book3:", book3)

	book6 := bookFunction(book3)
	book7 := bookFunction2(&book3)
	fmt.Println("book6, book7", book6, book7)

	fmt.Println("===============切片")
	var slice1 []int = make([]int, 5, 10)
	printSlice(slice1)

	slice1 = append(slice1, 1, 2)
	printSlice(slice1)

	// 切面截取：start:end => [0, 2)
	slice2 := slice1[:2]
	printSlice(slice2)

	// 切面截取：start:end => [1,len())
	slice3 := slice1[1:]
	printSlice(slice3)

	// 切面截取：start:end => [0, len())
	slice4 := slice1[:]
	printSlice(slice4)

	var i3 = 0
	var end = cap(slice1)
	for i3 = 0; i3 < end+1; i3++ {
		// slice1[i3] = i3
		slice1 = append(slice1, i3)
	}
	fmt.Println("append超过cap")
	printSlice(slice1)

	// 初始化长度 与 slice1 相同， 但 容量 扩大 2倍
	slice5 := make([]int, len(slice1), cap(slice1)*2)

	// 将 slice1 复制 给 slice5
	copy(slice5, slice1)
	printSlice(slice5)

	var slice6 []int
	if slice6 == nil {
		fmt.Println("slice6切片 是 空的")
		printSlice(slice6)
	}

	fmt.Println("=================range:")
	var hello = "hello"
	for i, v := range hello {
		fmt.Printf("range:字符串: index: %d, char: %c\n", i, v)
	}

	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println("range:数组", i, v)
	}

	for i, _ := range arr {
		fmt.Println("range: key: %s", i)
	}

	for _, v := range arr {
		fmt.Println("range: value: %s", v)
	}

	fmt.Println("=================map")

	// map用 make声明
	countryMap := make(map[string]string, 10)
	countryMap["1"] = "1"
	countryMap["2"] = "2"
	countryMap["3"] = "3"

	for key, value := range countryMap {
		fmt.Printf("key:%s, value:%s \n", key, value)
	}

	delete(countryMap, "1")

	for key, _ := range countryMap {
		fmt.Println("map删除后：", countryMap[key])
	}

	// map直接声明: 带参
	countryMap2 := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	fmt.Println("map size：", len(countryMap2))

	v1 := countryMap2["1"]
	v2, ok := countryMap2["2"]
	v3, ok2 := countryMap2["5"]
	fmt.Printf("v1:%s, v2:%s, v2Ok: %s, v3: %s, v3OK: %s \n", v1, v2, ok, v3, ok2)

	// map声明： 不带参
	countryMap3 := map[string]string{}
	for k, v := range countryMap2 {
		countryMap3[k] = v
	}

	for k, v := range countryMap3 {
		fmt.Printf("countryMap3: k:%s, v:%s \n", k, v)
	}

	fmt.Println("===============类型转换")
	intNum := 1
	// int 与 float互转
	floatNum := float32(intNum)
	intNum2 := int8(floatNum)
	fmt.Printf("类型转换：int => float:%s=>%s , fload => int: %s => %s \n", intNum, floatNum, floatNum, intNum2)

	// 字符串 与 int 互转
	stringNum := strconv.Itoa(intNum)
	intNum3, err := strconv.Atoi(stringNum)
	fmt.Printf("int => string: %s => %s, string => int: %s => %s, err: %s \n", intNum, stringNum, stringNum, intNum3, err)

	// 字符串 与 float互转
	stringNum2 := "3.14"
	floatNum2, _ := strconv.ParseFloat(stringNum2, 64)
	stringNum3 := strconv.FormatFloat(floatNum2, 'f', 4, 32)
	fmt.Printf("string => float: %s => %s, float => string: %s => %s \n", stringNum2, floatNum2, floatNum2, stringNum3)

	// 断言
	var stringInterface interface{} = "hello , go"
	strInterface2, ok := stringInterface.(string)

	if ok {
		fmt.Println("是string类型：", strInterface2)
	} else {
		fmt.Println("断言失败， 不是string类型")
	}

	// 空接口类型
	switchType(1)
	switchType("string喽")
	switchType(float64(123))

	// 结构体 转换
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, 结构体转换"

	// 打印 StringWriter 的字段值
	fmt.Println(sw)

	fmt.Println("================interface")
	var interfaceNum interface{} = 1
	// 类型断言
	interfaceType := interfaceNum.(int)
	fmt.Printf("type:%T \n", interfaceType)

	// 空接口
	interfaceFun("hello, interface")
	interfaceFun(123123)
	interfaceFun(12.0123)

	// 带参构造
	apple := FruitSize1{name: "apple"}
	orange := FruitSize2{name: "orange"}
	// new: 只能无参构造
	all := new(AllFruit)
	all2 := AllFruit{}

	fmt.Printf("多态：apple:action:%s, description:%s \n", apple.action(), apple.description())
	fmt.Printf("多态：orange:action:%s, description:%s \n", orange.action(), orange.description())
	fmt.Printf("多态：%s \n", all.action())
	fmt.Printf("多态：%s, \n", all2.action())

	var interfaceThings interface{}
	fmt.Printf("是否为空 ： %t \n", interfaceThings == nil)

	fmt.Println("==================泛型")
	// any约束
	xxPrint(1)
	xxPrint("hello")
	xxPrint(4.1123)
	// comparble约束： 类型只能适用于 == 或者 !=
	slice7 := []int{1, 2, 3}
	fmt.Println("泛型：compare约束：", compare(slice7, 3))
	// 联合约束
	fmt.Println("泛型：联合约束： 比较大小：", max2[int](1, 2))
	fmt.Println("泛型： 比较大小：", max2[float64](1.09, 1.10))

	// 自定义约束
	okString := OKStringer{"小红", 18}
	PrintThing(okString)

	fmt.Println("==================错误处理")

	// 显示调用：
	err = errors.New("这是一个显示调用的错误")
	fmt.Println(err)

	result, err2 := divide(2, 0)
	if err2 == nil {
		fmt.Println("result:", result)
	} else {
		fmt.Println("error:", err2)
	}

	// 自定义error
	fmt.Println(CustomError{404, "不中"}.error())

	// panic && revoer
	fmt.Println("1 panic")
	safeFunction()
	fmt.Println("2 panic")

	fmt.Println("=============goroutine")
	// goroutine的交互使用
	go sayHello()
	var i5 = 0
	for i5 = 0; i5 < 10; i5++ {
		fmt.Println("====go2:", i5)
		time.Sleep(100 * time.Millisecond)
	}

	// channle的使用： 其中 <- 代表：多个 goroutine之间可以进行 参数通信
	chan1 := make(chan int, 10)
	chan1 <- 1
	chan1 <- 2
	chan1 <- 3
	chan1 <- 4

	go fillChannel(chan1)
	for i := range chan1 {
		fmt.Println("接收chan:", i)
	}

	// select的使用：
	startChan := make(chan int, 5)
	endChan := make(chan int, 5)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 1; i++ {
			data := <-startChan
			fmt.Println("startChan: 接收：", data)
			wg.Done()
		}
		endChan <- 0
	}()

	// 此处 必定要用 wg的指针 ，否则 wg.add()是在副本中完成的， 对此处没有任何影响
	selectFunc(startChan, endChan, &wg)
	wg.Wait()
	fmt.Println("select执行结束： waitGorup")

	fmt.Println("================继承")
	// 组合
	dog := Dog{
		Animal: Animal{name: "动物"},
		run:    "跑起来了",
	}

	dog.call()
	fmt.Println("组合：run:", dog.run)

	// 接口
	var speaker Speaker
	dog = Dog{
		Animal: Animal{name: "动物"},
		run:    "跑起来了",
	}
	speaker = &dog
	speaker.call()
	fmt.Println("接口：run:", dog.run)

}

func sum2(num1, num2 int) int {
	return num1 + num2
}

func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
	// go 中 没有 三目运算符
	// return num1 > num2 ? num1 : num2
}

/**
* 返回多个参数
 */
func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func bookFunction(book Books) Books {
	book.author = "bookFunction"
	return book
}

func bookFunction2(book *Books) Books {
	book.author = "bookFunction指针引用"
	return *book
}

func printSlice(slice1 []int) {
	fmt.Println("slice1:", len((slice1)), cap(slice1), slice1)
}

func switchType(v interface{}) {
	switch v2 := v.(type) {
	case int:
		fmt.Println("int类型：", v2)
	case string:
		fmt.Println("string类型：", v2)
	default:
		fmt.Println("未知类型：", v2)
	}

}

// 定义一个接口 Writer
type Writer interface {
	write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

// 实现 Write 方法
// *StringWriter: 代表该方法 属于StringWriter
// 接收者变量名为 sw，在方法体内代表调用该方法的那个 StringWriter 实例
func (sw *StringWriter) write(data []byte) (int, error) {
	// 将输入的[]byte 转换为 string, 并用 += 的方式 完成拼接字符串
	sw.str += string(data)
	return len(data), nil
}

func interfaceFun(data interface{}) {
	fmt.Printf("data:%v, type: %T \n", data, data)
}

type Fruit interface {
	action() string
	description() string
}

type FruitSize1 struct {
	name string
}

type FruitSize2 struct {
	name string
}

type AllFruit struct{}

func (fs FruitSize1) action() string {
	result := "我爱吃" + fs.name
	return result
}

func (fs FruitSize1) description() string {
	return fs.name + "真大，真好吃"
}

func (fs FruitSize2) action() string {
	return "我也爱吃" + fs.name
}

func (fs FruitSize2) description() string {
	return fs.name + "真的特别大，特别好吃"
}

func (af AllFruit) action() string {
	return "所有水果我都爱吃"
}

func xxPrint[T any](data T) {
	fmt.Printf("data:%v, type: %T \n", data, data)
}

func compare[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// 联合约束
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func max2[T Number](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

type Stringer interface {
	string() string
}

type OKStringer struct {
	name string
	age  int
}

func (ok OKStringer) string() string {
	return fmt.Sprintf("name:%s, age:%v", ok.name, ok.age)
}

func PrintThing[T Stringer](param T) {
	fmt.Println(param.string())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("报错：被除数不能为0")
	}
	return a / b, nil
}

type Error interface {
	error() string
}

type CustomError struct {
	code        int
	description string
}

func (ce CustomError) error() string {
	return fmt.Sprint("自定义error: code: %v, description: %s \n", ce.code, ce.description)
}

func safeFunction() {
	// defer 函数： 延迟执行： 直到函数 返回前执行
	defer func() {
		// recover(): 判定当前的goroutine 是否遇到 panic, 如果遇到， 则捕获 panic， 并返回 panic
		// recover 只能在 defer函数中执行
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before panic")
	panic("something went wrong")
	// 不会执行
	fmt.Println("after panic")
}

func sayHello() {
	var i = 0
	for true {
		i++
		if i > 10 {
			break
		}
		fmt.Println(">>go1:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func fillChannel(c chan int) {
	var i int
	for i = 0; i < 5; i++ {
		c <- i
		fmt.Println("填充chan:", i)
	}

	// 如果 不关闭 chan, 最终会导致 所有goroutine处于 asleep状态， 最终 deadLock
	// 因为 后面的 range 会遍历 chan, 前提是 chan关闭了， 不然会一直等待chan的关闭
	close(c)
}

func selectFunc(startChan, endChan chan int, wg *sync.WaitGroup) {
	i := 0
	// select 语句使得一个 goroutine 可以等待多个通信操作。select 会阻塞，直到其中的某个 case 可以继续执行
	select {
	case startChan <- i:
		wg.Add(1)
		i++
		fmt.Println("startChan: 填充：", i)
	case <-endChan:
		fmt.Println("b 关闭了")
		return
	}

}

type Animal struct {
	name string
}

func (a Animal) call() {
	fmt.Println(a.name, "叫了 ！")
}

type Dog struct {
	Animal
	run string
}

type Speaker interface {
	call()
}
