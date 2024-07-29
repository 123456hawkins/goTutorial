package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type NewType int32

// 类型声明
// 一个新定义的类型和它的源类型为两个不同的类型。
// 一个新定义的类型和它的源类型的底层类型（将在下面介绍）一致并且它们的值可以相互显式转换。
type (
	NewType1 string //二者的底层类型都为string
	NewType2 bool
)
type Book struct {
	author string
	title  string
	pages  int
}

// 类型别名,表示同一类型
type (
	name = string
	age  = int
)

func main() {
	var n NewType = 1
	var m int32 = 1
	fmt.Println(reflect.TypeOf(n) == reflect.TypeOf(m))
	fmt.Println(unsafe.Sizeof(n))
}
