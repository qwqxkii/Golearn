package main

//在向该通道发送完数个整数后，检查通道是否已经关闭。如果没有关闭，则关闭它。
func main() {
	//编写一个函数，该函数接受一个整数通道作为参数。
	var ch1 = make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
	}()

}
