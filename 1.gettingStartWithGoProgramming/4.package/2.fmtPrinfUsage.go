package main

import "fmt"

//  `%v`：将被替换为对应实参字符串表示形式。
//  `%T`：将替换为对应实参的类型的字符串表示形式。
//  `%x`：将替换为对应实参的十六进制表示。实参的类型可以为字符串、整数、整数数组（array）或者整数切片（slice）等。 （数组和切片将在以后的文章中讲解。）
//  `%s`：将被替换为对应实参的字符串表示形式。实参的类型必须为字符串或者字节切片（byte slice）类型。
//  `%%`：将被替换为一个百分号。
func main() {
	a, b := 123, "go"
	fmt.Printf("数值为%v,字符串为%s", a, b)
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)
	fmt.Printf("1%% 50%% 99%%\n")
}
