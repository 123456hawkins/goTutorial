package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) //随机睡眠0到2.5s
	}
	wg.Done() //通知当前任务已完成
}
func main() {
	// 日志输出的标志设置为0，这意味着不添加任何额外的时间、日期或文件信息等。
	log.SetFlags(0)
	wg.Add(2) //注册两个新任务
	go SayGreetings("hi", 10)
	go SayGreetings("hello", 10)
	wg.Wait() //阻塞这里，等待所有任务都已完成
}
