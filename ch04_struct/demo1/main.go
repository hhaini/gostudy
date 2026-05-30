package main

import "fmt"

type Book struct {
	Title   string         // 书名
	Pages   int            // 书的页数
	Indexes map[string]int // 书的索引
}

func main() {
	var book = Book{Title: "The Go Programming Language", Pages: 700, Indexes: make(map[string]int)}
	fmt.Println(book.Pages, book.Title, book.Indexes)
}
