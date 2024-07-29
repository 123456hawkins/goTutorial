// 结构体
package main

import (
	"fmt"
	"reflect"
)

// 无名结构体
type Book1 struct {
	title, location string
	author          string
	pages           int
}

func main() {
	book := Book1{author: "feawefaerg", title: "feawegae", pages: 20, location: "23324234"}
	book2 := Book1{}
	fmt.Println(book)
	fmt.Println(book2) //默认值为"" "" 0
	fmt.Println(reflect.TypeOf(book))
	book2 = book
	book = Book1{author: "feawefae123rg", title: "feaweg123ae", pages: 200, location: "23324234"} //必须要加Book1,否则赋值错误
	v := &book
	fmt.Print(v.author) //v.author等价为(*v.author),go会默认解引用
	fmt.Println(&v.author == &book.author)
}
