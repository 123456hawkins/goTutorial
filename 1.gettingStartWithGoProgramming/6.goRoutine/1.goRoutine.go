package main

import (
	"log"
	"math/rand"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) //随机睡眠0到2.5s
	}
}
func main() {
	log.SetFlags(0)
	go SayGreetings("hi", 10)
	go SayGreetings("hello", 10)
	time.Sleep(5 * time.Second)
}
