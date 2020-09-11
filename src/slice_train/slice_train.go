package main

import "fmt"

func main() {
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	// 打印原始切片
	fmt.Println("numbers ==", numbers)
	// 打印子切片从索引1(包含)到索引4(不包含)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	// 默认下限为0
	fmt.Println("numbers[:3] ==", numbers[:3])
	// 默认上限为len(s)
	fmt.Println("numbers[4:] ==", numbers[4:])
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	// 打印子切片从索引0(包含)到索引2(不包含)
	number2 := numbers[:2]
	printSlice(number2)
	// 打印子切片从索引2(包含)到索引5(不包含)
	number3 := numbers[2:5]
	printSlice(number3)

	var newNumbers []int
	printSlice(newNumbers)
	// 允许追加空切片
	newNumbers = append(newNumbers, 0)
	printSlice(newNumbers)
	// 向切片添加一个元素
	newNumbers = append(newNumbers, 1)
	printSlice(newNumbers)
	// 同时添加多个元素
	newNumbers = append(newNumbers, 2, 3, 4)
	printSlice(newNumbers)
	// 创建切片newNumbers1是之前切片的两倍容量
	newNumbers1 := make([]int, len(newNumbers), (cap(newNumbers))*2)
	// 拷贝newNumbers的内容到newNumbers1
	copy(newNumbers1, newNumbers)
	printSlice(newNumbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
