package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("下一个伪随机数为", rand.Uint32())
	fmt.Println(time.Now())
}
