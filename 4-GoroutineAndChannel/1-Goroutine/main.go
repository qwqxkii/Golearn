package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		fmt.Println("one")
	}()

	go func() {
		fmt.Println("two")
	}()
	go func() {
		fmt.Println("three")
	}()
	go func() {
		fmt.Println("four")
	}()
	go func() {
		fmt.Println("five")
	}()
	//获取并发线程数量 #
	fmt.Print(runtime.GOMAXPROCS(0))
	time.Sleep(1 * time.Second)

}
