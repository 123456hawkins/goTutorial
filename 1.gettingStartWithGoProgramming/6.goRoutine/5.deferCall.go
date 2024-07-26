package main

import (
	"fmt"
	"time"
)

func fnDefer1() {
	defer fmt.Println("first LINE")
	defer fmt.Println("second Line")
	fmt.Println("third line")
}
func fnDefer2() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
}

// 一个延迟调用可以修改包含此延迟调用的最内层函数的返回值
func fnDefer3(n int) (r int) {
	defer func() {
		r += n //修改返回值
	}()
	return n + n
}

// 延迟调用实参估值
func fnDeterActualParamsEvaluate() {
	fmt.Println("----paramsEvaluate-------")
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()
	fmt.Println()
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer func(i int) {
				fmt.Println("b:", i+x)
			}(i)
		}
		x = 10
	}()
}

// 协程调用估值时刻
func fnGoRoutineEval() {
	fmt.Println("-----fnGoRoutineEval----")
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 123 789
	}(a) //在被定义的时候就已经确定了a=123
	a = 789
	time.Sleep(2 * time.Second)
}
func main() {
	fnDefer1()
	fnDefer2()
	fmt.Println(fnDefer3(5)) //15
	fnDeterActualParamsEvaluate()
	fnGoRoutineEval()
}
