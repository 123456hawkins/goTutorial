// 匿名函数
package main
func main(){
	x,y:=func()(int, int){
		return 128,128
	}()//注意结尾的()

	e,f:=func(a,b int)(e,f int){
		 e=a*b
		 f=a-b
		 return
	}(22,333)
	
	println(x,y)
	println(e,f)
}