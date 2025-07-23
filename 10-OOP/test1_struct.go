package main

import "fmt"

// 声明一种行的数据类型 myint， 是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

type myString string

func changeBook(book Book) {
	//传递一个book的副本
	fmt.Printf("book type is %T\n", book)
	fmt.Printf("book type is %p\n", &book)
	fmt.Printf("book type is %T\n", book.auth)
	fmt.Printf("book type is %p\n", &book.auth)
	book.auth = "666"
}

func changeBook2(book *Book) {
	//指针传递
	fmt.Printf("book type is %v\n", book)
	fmt.Printf("book type is %T\n", book)
	fmt.Printf("book type is %p\n", &book)
	fmt.Printf("book type is %v\n", book.auth)
	fmt.Printf("book type is %p\n", &book.auth)
	book.auth = "777"
}

func main() {

	/*	var a myint = 10
		fmt.Println("a = ", a)
		fmt.Printf("type of a = %T\n", a)*/
	/*	var b myString = "hello"
		fmt.Println(b)
		fmt.Printf("type of a = %T\n", b)*/
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)
	fmt.Printf("%p\n", &book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
