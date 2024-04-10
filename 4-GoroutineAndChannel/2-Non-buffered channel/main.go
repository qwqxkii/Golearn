package main

import "fmt"

//Non-buffered channel

func main() {
	//非缓冲通道
	var ch = make(chan string)
	//缓冲通道
	//var ch = make(chan string,10)
	go func() {
		//不在go fun中运行就会出现死锁
		ch <- "hello world!"
	}()

	h := <-ch
	fmt.Println(h)
}
