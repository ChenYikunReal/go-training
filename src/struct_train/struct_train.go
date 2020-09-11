package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"疯狂Java讲义", "李刚", "Java教程", 123456})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "疯狂Python讲义", author: "李刚", subject: "Python教程", bookId: 654321})
	// 忽略的字段为0或空
	fmt.Println(Books{title: "疯狂XML讲义", author: "李刚"})

	var Book1 Books
	var Book2 Books
	Book1.title = "疯狂Android讲义"
	Book1.author = "李刚"
	Book1.subject = "Android教程"
	Book1.bookId = 234567
	Book2.title = "疯狂Swift讲义"
	Book2.author = "李刚"
	Book2.subject = "Swift教程"
	Book2.bookId = 765432
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookId)

	// 结构体指针
	printBook(&Book1)
	printBook(&Book2)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
