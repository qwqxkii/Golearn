package main

import "fmt"

func main() {
	demo1()
	demo2()
	demo3()
}
func demo1() {
	x := 100
	if x == 100 {
		fmt.Println("好")
	}
}
func demo2() {
	n := 100
	if n > 10 {
		println("nb")
	} else {
		println("easy")
	}
}
func demo3() {
	if x := 100; x > 10 {
		println("n1")
	} else if x > 30 {
		println("n2")
	} else {
		println("easy")
	}
}
