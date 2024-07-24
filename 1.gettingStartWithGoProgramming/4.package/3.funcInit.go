package main

import (
	"fmt"
	random "math/rand" //包别名
	_ "net/http/pprof" //匿名引入,会执行相关包的init函数
	. "time"
)

var abc string

func init() {
	abc = "dafawfe"
}
func main() {
	fmt.Printf(abc)
	fmt.Print(random.Uint32())
	fmt.Print(Now())
}
