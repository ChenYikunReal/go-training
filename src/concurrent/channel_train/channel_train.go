package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 把sum发送到通道c
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据
	// 从通道c中接收
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
