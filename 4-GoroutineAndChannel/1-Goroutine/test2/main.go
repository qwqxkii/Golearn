package main

import (
	"fmt"
	"time"
)

// 使用Goroutine和Channel通信
// 创建两个Goroutine：，
// 一个发送当前时间的字符串到一个Channel
func send(str string) {
	ch <- str
}

// 另一个从该Channel接收这个字符串并打印出来。
func get() {
	qwq := <-ch
	fmt.Println(qwq)
}

var ch = make(chan string, 2)

func main() {
	//创建两个Goroutine
	go send("20240410")
	go get()

	time.Sleep(1 * time.Second)

}
