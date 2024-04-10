package main

import (
	"fmt"
	"time"
)

// 创建和启动一个Goroutine
// 编写一个Go程序，在一个新的Goroutine中打印出"Hello, Goroutine!"。

func main() {
	go func() {
		fmt.Println("Hello, Goroutine!")
	}()
	time.Sleep(1 * time.Second)
}
