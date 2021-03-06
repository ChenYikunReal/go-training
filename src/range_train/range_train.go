package main

import "fmt"

func main() {
	// 这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	// 使用_不需要知道索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 使用i存储索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// range也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
