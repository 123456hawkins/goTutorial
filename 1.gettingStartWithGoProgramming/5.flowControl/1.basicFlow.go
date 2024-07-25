package main

import (
	"fmt"
	"math/rand"
)

func testIfElse(a, b int) string {
	if a > b {
		return "a比b大"
	} else if a < b {
		return "a比b小"
	} else if a == b {
		return "a等于b"
	} else {
		return ""
	}
}

// ifelse标准写法
func testIfElseStandard() {
	// 初始化代码;条件
	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "是一个偶数")
	} else {
		fmt.Println(n)
	}
}
func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i := 10; ; i++ {

		if i > 20 {
			break
		}
		fmt.Println(i)
	}
}
func testFor2() {
	for i := 20; i < 30; i++ {
		fmt.Print(i) //打印外部变量i
		i := i
		i = 10
		fmt.Println(i) //打印局部变量i
	}
}
func testContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
func testRange() {
	fmt.Println("------------")
	for i := range 10 {
		fmt.Println(i)
	}
}
func testSwitch(testNum int64) {
	switch testNum {
	case 1:
		fmt.Println("awefawe")
	case 2:
		fmt.Println("fefergfw")
	case 3:
		fmt.Println("11gere")
	default:
		fmt.Println("default")

	}
}
func testSwitch2() {
	switch n := rand.Intn(10000); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		break // 这里的break语句可有可无的，效果
		// 是一样的。执行不会跳到下一个分支。
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
	// 上一行可能编译不过，因为6和上一个case中的
	// 6重复了。是否能编译通过取决于具体编译器实现。
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")

	}
}
func testFallThrough() {
	fmt.Println("-----testFallThrough----------")
	switch n := rand.Intn(1000) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n=", n)
		fallthrough
	case 5, 6, 7, 8:
		n := 99 //新声明的n,仅在这个分支可见
		fmt.Println("n=", n)
		fallthrough
	default:
		fmt.Println("n=", n)
	}

}
func testGoTo() {
	fmt.Println("----testGoTo-----")
	i := 0
Next:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}
}
func testGoTo2() {
	fmt.Println("----testGoTo2------")
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}
	// 创建一个显式代码块以缩小k的作用域。
	// 如果没有代码块就会报错,因为exit标签声明在k的作用域内，但它的使用在k的作用域之外。
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next
Exit:
	fmt.Println("exit")
}

// break和continue结合跳转标使用
func testGoTo3(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer //跳出流程控制代码块
			case n%i == 0:
				continue Outer //继续流程控制代码块
			}
		}
	}
	return n
}
func main() {
	fmt.Println(testIfElse(2, 3))
	testIfElseStandard()
	testFor()
	testFor2()
	testContinue()
	testRange()
	testSwitch(50)
	testSwitch2()
	testFallThrough()
	testGoTo()
	testGoTo2()
	k := testGoTo3(266)
	fmt.Println(k)
}
