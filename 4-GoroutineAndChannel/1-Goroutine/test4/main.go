package main

import (
	"fmt"
	"runtime"
	"sync"
)

// //设置GOMAXPROCS并观察行为变化
// 编写一个Go程序，，尝试更改runtime.GOMAXPROCS()的值，并观察和记录程序的行为变化。

func main() {
	//它并发地启动一定数量的Goroutine
	var sw sync.WaitGroup
	sw.Add(5)
	for i := 1; i < 6; i++ {

		fmt.Println("当前GOMAXPROCS：", runtime.GOMAXPROCS(0))
		if i == 3 {
			fmt.Println("设置大小")
			runtime.GOMAXPROCS(2)
			fmt.Println("当前GOMAXPROCS：", runtime.GOMAXPROCS(0))
		}
		go func(i int) {
			defer sw.Done()
			//每个Goroutine都打印出其编号。
			println("编号:", i)
		}(i)

	}
	sw.Wait()

}
