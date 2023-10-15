package main

import "fmt"

type Book struct {
	ISBN   int
	Title  string
	Author string
}

func main() {
	var book Book
	book.ISBN = 43214
	book.Title = "go"
	book.Author = "xiaxuan"
	fmt.Println(book)
	fmt.Println(book.Title)
}
