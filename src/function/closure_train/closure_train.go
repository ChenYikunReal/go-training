package main

import "fmt"

// 该函数的目的是在闭包中递增i变量
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	num1 := getSequence()
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())
	num2 := getSequence()
	fmt.Println(num2())
	fmt.Println(num2())
}
