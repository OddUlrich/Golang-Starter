package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var book1 Book
	var book2 Book

	book1.title = "新华字典"
	book1.author = "人民出版社"
	book1.subject = "语言工具书"
	book1.bookId = 64953407

	book2.title = "左岸"
	book2.author = "文学出版社"
	book2.subject = "小说"
	book2.bookId = 64957040

	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.bookId)

	fmt.Printf("Book 2 title : %s\n", book2.title)
	fmt.Printf("Book 2 author : %s\n", book2.author)
	fmt.Printf("Book 2 subject : %s\n", book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", book2.bookId)
}
