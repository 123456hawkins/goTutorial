package main

func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}

// 匿名返回结果
func SquareOfSumAnonymous(a int64, b int64) (int64, int64) {
	x, y := a+b, a-b
	s := x * x
	d := y * y
	return s, d
}

// go不支持输入参数默认值
func DefaultParam() (x int, y bool) {
	println(x, y)
	return
}

// 若干联系相同的参数类型相同可以只写一个
func SquaresOfSumAndDiffSameType(a, b, c, d int64) (j, k int64) {
	return (a + b + c + d), (a * b * c * d)
	// 等价于j=(a+b+c+d);k= (a*b*c*d); return
}

// 如果只有一个返回结果，并且返回参数是匿名的，可以不用使用_符号
func fnAnonymousReturnParams() string {
	return "hahahah"
}
func main() {
	s, d := SquaresOfSumAndDiff(100, 100)
	println(s, d)
	a, b := SquareOfSumAnonymous(200, 100)
	println(a, b)
	DefaultParam()
	e, f := SquaresOfSumAndDiffSameType(1, 2, 3, 4)
	println(e, f)
	println(fnAnonymousReturnParams())
}
