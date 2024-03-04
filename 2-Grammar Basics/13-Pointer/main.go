package main

import "fmt"

// 指针
// 指针的值是一个变量地址 ：取地址符&     *根据地址取值
func main() {
	getadress()
	demo1()
}
func getadress() {
	x := 100
	p := &x
	fmt.Println("%T, %v", p, p)
}
func demo1() {
	x := 10
	fmt.Printf("%p\n", &x)

	var p *float64
	qwq := &p
	fmt.Printf("%p\n", qwq)

}

func demo2() {

}
