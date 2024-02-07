package main

import "fmt"

type BookPointer struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var book1 BookPointer /* 声明 book1 为 BookPointer 类型 */
	var book2 BookPointer /* 声明 book2 为 BookPointer 类型 */

	book1.title = "新华字典"
	book1.author = "人民出版社"
	book1.subject = "语言工具书"
	book1.bookId = 64953407

	book2.title = "左岸"
	book2.author = "文学出版社"
	book2.subject = "小说"
	book2.bookId = 64957040

	/* 打印 book1 信息 */
	printBookPointer(&book1)

	/* 打印 book2 信息 */
	printBookPointer(&book2)
}
func printBookPointer(book *BookPointer) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
