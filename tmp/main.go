package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，该程序启动若干Goroutine，
// 主函数应该从这个channel中读取值直到所有Goroutine都发送完毕。使用sync.WaitGroup来同步，并确保所有数据都被接收。
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	var ch = make(chan int, 5)

	//// 每个Goroutine
	for i := 1; i < 6; i++ {
		go func(i int) {
			defer wg.Done()
			// 向一个共享的chan int类型的channel发送一个唯一的整数。
			ch <- i
			println("已传入:", i)

		}(i)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		fmt.Println("已接收", i)
	}

}
