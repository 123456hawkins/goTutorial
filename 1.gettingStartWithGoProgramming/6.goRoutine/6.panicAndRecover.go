package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("hi")
	defer func() {
		v := recover()
		fmt.Println("恐慌恢复了", v)
	}()
	panic("byebye")
	// fmt.Println("cannont excutes")
}
