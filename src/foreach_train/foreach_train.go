package main

import "fmt"

func main() {
	strings := []string{"A", "B"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位x的值是 %d\n", i, x)
	}
}
