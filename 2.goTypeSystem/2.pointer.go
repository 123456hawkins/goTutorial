package main

import (
	"fmt"
	"reflect"
)

//	func testPointer(*data int) {
//		fmt.Println(data)
//	}
type Ptr *int //Ptr是具名指针类型，基类型为int
type PP **int //PP是具名多级指针类型，基类型为Ptr

func doubleNum(x *int) *int {
	// 如果直接返回int则会直接分配新的地址
	*x += *x
	*x++ //指针不能增减默认转换为(*x)++
	return x
}
func main() {
	num := 2
	numPointer := &num
	numPointer2 := &numPointer
	fmt.Println(numPointer)
	fmt.Println(reflect.TypeOf(numPointer)) //

	fmt.Println(numPointer2)
	fmt.Println(reflect.TypeOf(numPointer2)) //指针的指针

	v := new(int32) //分配一块新的地址空间
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))

	fmt.Println(*numPointer) //解引用

	a := 3
	b := doubleNum(&a) //直接修改a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b, *b, &a == b)
}
