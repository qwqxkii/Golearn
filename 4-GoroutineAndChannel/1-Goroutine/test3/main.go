package main

import (
	"math/rand"
	"sync"
	"time"
)

//3使用WaitGroup等待Goroutines完成

//主程序应使用sync.WaitGroup来等待所有Goroutine完成后再退出。

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	//编写一个Go程序，启动三个Goroutine，
	for i := 1; i < 4; i++ {
		////每个都在等待随机的时间后打印出其编号（1、2或3）。
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3-1+1)+1) * time.Second)
			print(i)
		}(i)
	}
	wg.Wait()
}
