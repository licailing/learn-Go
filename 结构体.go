package main

import "fmt"

/*
结构体可以存储不同数据类型
结构体表示一项纪录
*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	/* 声明Book1为Books类型 */
	var Book1 Books

	/* 访问结构体成员 */
	Book1.title = "Go 语言"
	Book1.author = "www.shouce.ren"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	printBook(Book1)
	/*
		Book title:Go 语言
		Book author:www.shouce.ren
		Book subject:Go 语言教程
		Book book_id:6495407
	*/

	printBook1(&Book1) //传址
	/*
		Book title:Go 语言
		Book author:www.shouce.ren
		Book subject:Go 语言教程
		Book book_id:6495407
	*/
}

/* 向函数传递结构体 */
func printBook(Book Books) {
	fmt.Printf("Book title:%s\n", Book.title)
	fmt.Printf("Book author:%s\n", Book.author)
	fmt.Printf("Book subject:%s\n", Book.subject)
	fmt.Printf("Book book_id:%d\n", Book.book_id)
}

/* 结构体指针 */
func printBook1(Book *Books) {
	fmt.Printf("Book title:%s\n", Book.title) //结构体指针访问结构体成员
	fmt.Printf("Book author:%s\n", Book.author)
	fmt.Printf("Book subject:%s\n", Book.subject)
	fmt.Printf("Book book_id:%d\n", Book.book_id)
}
