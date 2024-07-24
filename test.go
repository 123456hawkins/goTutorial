package main
import 'math/rand'
const MaxRand=16
func StatRandomNumbers(numRands int)(int,int){
	var a,b int
	for i := 0; i < numRands; i++ {
		if rand.Intn(MaxRand)<MaxRand/2 {
			a=a+1
		}else{
			b++
		}
	}
	return a,b
}
func main()  {
	var num=100
	x,y=:=StatRandomNumbers(num)
	print('result',x,'+',y,'=',num,"?")
	print(x+y==num)
}