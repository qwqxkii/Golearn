package main

import "fmt"

func main() {
	panicdemo()
}

// 主动出触发
func panicdemo() {
	panic("出现致命错误，请联系开发者")
}

// 除0
func panicdemo2() {
	fmt.Println("不能除以0")
	n := 5
	println(n / 0)
	//# learn/3-TypeOfData/14-panic
	//	.\main.go:18:14: invalid operation: division by zero
}
