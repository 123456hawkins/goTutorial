package main

import "fmt"

const Size = 32

type Person struct {
	name string
	age  int
}
type Language struct {
	name string
	year int
}

// 用容器和nil值比较判断是否为0值
func compareContainer() {
	fmt.Println("----compareContainer-----")

	var s []int
	var m map[string]int
	fmt.Println(s == nil, m == nil)

}

// 获取容器长度
func getContainerLen() {
	fmt.Println("----getContainerLen-----")
	// cap是容量，length是数组元素
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6, 7}
	c := map[int]string{1: "efawef", 2: "feawef"}
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(len(c)) //map无法使用cap
}

// 数组赋值
func arrAssign() {
	fmt.Println("----arrAssign-----")
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a2 := a1
	a1[0] = 2
	fmt.Println(a1, a2) //[2 2 3 4 5 6 7 8 9 0] [1 2 3 4 5 6 7 8 9 0]

	// map共用同一引用
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := m1
	m1["a"] = 2
	fmt.Println(m1, m2) //map[a:2 b:2 c:3] map[a:2 b:2 c:3]

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1[0] = 2
	fmt.Println(s1, s2) //[2 2 3 4 5] [2 2 3 4 5]
}

// 添加或者删除容器
func addOrDel() {
	fmt.Println("----addOrDel-----")
	// map
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
	m["e"] = 5
	fmt.Println(m)
	// 切片
	s := []int{1, 2, 3}
	s1 := append(s, 0) //添加一个元素
	fmt.Println(s1)
	s3 := append(s, s1...) // 以s为基础添加s1中所有的元素,...类似于js中的数组解构
	fmt.Println(s3)
	s4 := append(s3) //等价于s4:=s3
	fmt.Println(s4)
	// 使用make来创建切片
	s5 := make([]int, 3, 5)
	fmt.Println(s5) //[0 0 0]

}

// 修改map值,如果元素类型为数组或者结构体类型,必须整体修改
func changeMapVal() {
	fmt.Println("----changeMapVal-----")
	type T struct {
		name string
		age  int
	}
	m1 := map[string]T{"A": {"FWEFWE", 12}}
	// m1["A"].age=18//error
	fmt.Println(m1)
	m2 := map[string][5]int{"a": {1, 2, 3, 4, 5}, "b": {2, 3, 4, 5, 6}}
	// m2["a"][1] = 2//error
	fmt.Println(m2)

	// 正确做法,先赋值给临时变量,然后再用临时变量整体覆盖要修改的映射元素
	t := m1["A"]
	t.age = 18
	m1["A"] = t
	fmt.Println(m1)

	a := m2["a"]
	a[1] = 999
	m2["a"] = a
	fmt.Println(m2)
}

// 使用切片语法
func useSlice() {
	fmt.Println("----useSlice-----")
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := a[:] //s1=[0:9]
	s2 := s1[2:6]
	s3 := s1[1:]
	s4 := s1[3:5]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
}
func sliceToArr() {
	fmt.Println("----sliceToArr-----")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a := [3]int(s[0:3])
	fmt.Println(a, len(a), cap(a))
}
func useSliceCopy() {
	fmt.Println("----useSliceCopy-----")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := []int{41, 51, 61, 1, 2, 3, 4, 5}
	n := copy(s, s1)      //s1复制给s,并返回复制了多少个元素n
	fmt.Println(n, s, s1) //8 [41 51 61 1 2 3 4 5 9 0] [41 51 61 1 2 3 4 5]
}
func useForRange() {
	fmt.Println("----useForRange-----")
	type T struct {
		age  int
		name string
	}
	m1 := map[string]T{"A": {12, "WEFWEF"}, "B": {121, "WEFWEF2232"}}
	for key, element := range m1 { //被遍历的容器是m1的副本
		fmt.Println(key, element) //遍历顺序是随机的
	}
	for key, _ := range m1 {
		delete(m1, key)
	}
	fmt.Println(m1)

	a1 := [...]int{1, 2, 3, 4, 5, 6}
	for index, element := range a1 { //此遍历用的是a1中的一个副本
		a1[index] = 6
		fmt.Println(element) //遍历的element不和原数组共享数据,所以修改a1中的值,遍历的element的值不会修改,
		element = 22         //修改element不会影响原数组
	}
	fmt.Println(a1)
	for i, _ := range a1 { //当遍历较大的数组时,推荐舍弃掉后面的element来提高性能
		fmt.Println(a1[i])

	}
}
func usePointerArrAsArr() {
	fmt.Println("----useForRange-----")
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	for i, v := range &a1 { //复制指针的开销很小
		fmt.Println(i, v)
	}
	for _, v := range a1[:] { //复制切片开销很小
		fmt.Println(v)
	}
	p := &a1
	fmt.Println(p[1], p[2])
}
func useClear() {
	fmt.Println("----useForRange-----")
	// 用于清理切片和map
	a1 := []int{1, 2, 3, 4, 5, 6}
	clear(a1)
	fmt.Println(a1)
	a2 := []int{1, 2, 3, 4, 5, 6}
	clear(a2[2:5])
	fmt.Println(a2)
}
func moreSliceOperate() {
	fmt.Println("----moreSliceOperate-----")
	a1 := []int{1, 2, 3, 4, 5, 6}
	a1clone := append(a1[:0:0], a1...)
	fmt.Println(a1clone)
	// 推荐以下写法
	if a1 != nil {
		a1Clone2 := make([]int, len(a1))
		copy(a1Clone2, a1)
	}
	fmt.Println(a1clone)

	// 删除切片元素
	// 第一种写法
	s1 := []int{1, 2, 3, 4, 5, 6}
	s1 = append(s1[:3], s1[4:]...) //删除第四个元素
	fmt.Println(s1)
	// 第二种写法
	s1 = s1[:3+copy(s1[3:], s1[4:])]
	fmt.Println(s1)

	// 将一个切片的所有元素插到另一个切片中
	// 低效写法
	s0 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{99, 88, 77}
	s0 = append(s0[:2], append(s2, s0[2:]...)...) //开辟了两次内存空间
	fmt.Println(s0)
	// 更高效写法
	ss0 := []int{1, 2, 3, 4, 5}
	ss1 := []int{99, 77, 66}
	ss0 = insertSliceFn(ss0, ss1, 5)
	fmt.Println(ss0)
	// 前弹出
	cs1, e1 := ss1[1:], ss1[0]
	fmt.Println(cs1, e1)
	// 后弹出
	cs2, e2 := ss1[:len(ss1)-1], ss1[len(ss1)-1]
	fmt.Println(cs2, e2)
	// 前入队
	ss1 = append([]int{2, 3, 4, 5}, ss1...)
	fmt.Println(ss1)
	// 后入队
	ss1 = append(ss1, []int{2, 3, 4, 5}...)
	fmt.Println(ss1)
}
func insertSliceFn(originSlice []int, insertSlice []int, insertPlace int) []int {
	if insertPlace <= len(originSlice) && insertPlace >= 0 {
		if insertPlace == len(originSlice) {
			return append(originSlice, insertSlice...)
		} else if insertPlace == 0 {
			return append(insertSlice, originSlice...)
		} else {
			if cap(originSlice) >= len(originSlice)+len(insertSlice) {
				originSlice = originSlice[:len(originSlice)+len(insertSlice)]
				copy(originSlice[insertPlace+len(insertSlice):], originSlice[insertPlace:])
				copy(originSlice[insertPlace:], insertSlice)
			} else {
				x := make([]int, 0, len(insertSlice)+len(originSlice))
				x = append(x, originSlice[:insertPlace]...)
				x = append(x, insertSlice...)
				x = append(x, originSlice[insertPlace:]...)
				originSlice = x
			}
			return originSlice
		}
	} else {
		return nil
	}
}
func main() {
	// 数组
	x := [6]Person{{name: "feawef", age: 2342}}
	fmt.Println(x, x[0].age)
	// 切片
	y := []Person{{name: "12312", age: 111}}
	fmt.Println(y, y[0].age)
	// map
	z := map[string]Person{"a": {name: "12312", age: 12312}}
	fmt.Println(z, z["a"])
	// ...戴表自动推断数组长度
	w := [...]int32{12, 32, 4654, 784, 1454, 65}
	fmt.Println(len(w))
	// fmt.Println(&z["a"]) //报错,不能直接对map寻址
	p := z["a"]
	fmt.Println(&p) //间接取址

	q := [...]Language{{"awefawe", 123343}, {"faeegre", 34234}} //简写
	q1 := map[string]Language{"a": {"efafwef", 123123}, "b": {"efwefwe", 5555}}
	fmt.Println(q, q1)
	compareContainer()
	getContainerLen()
	arrAssign()
	addOrDel()
	changeMapVal()
	useSlice()
	sliceToArr()
	useSliceCopy()
	useForRange()
	usePointerArrAsArr()
	useClear()
	moreSliceOperate()
}
