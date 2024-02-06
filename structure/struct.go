package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {

	fmt.Println(Books{"新华字典", "人民出版社", "语言工具书", 64954207})

	fmt.Println(Books{title: "新华字典", author: "人民出版社", subject: "语言工具书", bookId: 64915407})

	fmt.Println(Books{title: "新华字典", author: "人民出版社"})
}
