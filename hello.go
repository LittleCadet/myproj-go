package main // 这是必须的，声明包名

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

var age string

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
