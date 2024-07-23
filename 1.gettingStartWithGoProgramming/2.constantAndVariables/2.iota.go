package main

const (
	Failed=iota-1//-1
	Unknown//0
	Succeed//1
)
const (
	Readable=1<<iota//1
	Writable//2
	Executable//4
)

func main(){
	const(
		k=3
		m float32 = iota+.5//1+.5
		n//2+.5
		o//3+.5
		p//4+.5
		q=iota*2//5*2
	)
	println(k,m,n,o,p,q)
} 