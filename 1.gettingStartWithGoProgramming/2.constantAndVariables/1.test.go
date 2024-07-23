package main
// 声明常量
const Π=3.1415926
const PI=Π
const x float32=3.14
// const a int8 = 256//报错

const (
	i float32=3.14
	j //自动补全为float32=3.14
	k //自动补全为float32=3.14

	A, B = "Go", "language"
	C,_//自动补全为"Go","language"
)

func main(){
	// 类型转换
	println(string(65))
	println(string(-1))
	// println(string(65.123))//报错
	// println(string(true))
	println(PI)
	println(x)
	println(i,j,k)
	println(C)
}