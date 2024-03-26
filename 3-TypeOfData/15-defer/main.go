package main

import "fmt"

func main() {
	defer A()
	B()
}
func A() {
	defer fmt.Println("A执行完毕")
	println("A开始执行")
}
func B() {
	defer fmt.Println("B执行完毕")
	println("B开始执行")
}
