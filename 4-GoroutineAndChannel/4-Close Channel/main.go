package main

import "fmt"

func main() {
	//关闭通道
	var ch1 = make(chan int)
	close(ch1)
	fmt.Println("通道已关闭")
}
