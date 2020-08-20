package main

import "fmt"

func main() {
	var var1 = "test"
	// test
	fmt.Println(var1)
	var var2 int
	// 0
	fmt.Println(var2)
	var var3 bool
	// false
	fmt.Println(var3)
	var var4 *int
	// <nil>
	fmt.Println(var4)
	var var5 []int
	// []
	fmt.Println(var5)
	var var6 map[string]int
	// map[]
	fmt.Println(var6)
	var var7 chan int
	// <nil>
	fmt.Println(var7)
	var var8 func(string) int
	// <nil>
	fmt.Println(var8)
	var var9 error
	// <nil>
	fmt.Println(var9)
	var var10 string
	//
	fmt.Println(var10)
	var var11 float64
	// 0
	fmt.Println(var11)
	// 它只能被用在函数体内，而不可以用于全局变量的声明与赋值
	var12, var13 := 1, 2.3
	// 1
	fmt.Println(var12)
	// 2.3
	fmt.Println(var13)
	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
		var14 = 1
		var15 = "Hello"
	)
	// 1
	fmt.Println(var14)
	// Hello
	fmt.Println(var15)
	// 直接定义此行的统一类型
	var var16, var17 int
	var16, var17 = 16, 17
	// 可以直接交换
	var16, var17 = var17, var16
	// 17
	fmt.Println(var16)
	// 16
	fmt.Println(var17)
	// _表示放弃值 是一个只写变量
	var _ = 1
	// 可以写入
	_ = 2
}
