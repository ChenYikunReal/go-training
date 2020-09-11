package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	const a, b, c = 1, false, "str" //多重赋值

	var area int = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println()
	fmt.Println(a, b, c)

	const (
		d = "def"
		e = len(d)
		f = unsafe.Sizeof(d)
	)
	fmt.Println(d, e, f)

	const (
		const0 = iota // 0
		const1        // 1
		const2        // 2
	)
	fmt.Println(const0, const1, const2)

	const (
		constant1 = iota // 0
		constant2        // 1
		constant3        // 2
		constant4 = "ha" // 独立值，iota += 1
		constant5        // "ha"   iota += 1
		constant6 = 100  // iota +=1
		constant7        // 100  iota +=1
		constant8 = iota // 7,恢复计数
		constant9        // 8
	)
	fmt.Println(constant1, constant2, constant3, constant4, constant5, constant6, constant7, constant8, constant9)

	// 使用iota的位移运算
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Printf("i=%d, j=%d, k=%d, l=%d\n", i, j, k, l)
}
