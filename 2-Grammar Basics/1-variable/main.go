package main

import "fmt"

func main() {
	var num int = 200
	println(num)

	//直接定义
	qwq := "201"
	println(qwq)

	//同时定义多个变量
	var (
		name   = "中文"
		age    = 25
		gender = "man"
	)

	println(name, age, gender)

	var q1 = 222
	var q2 int = 333
	fmt.Println("隐式类型定义", q1, "显示类型定义", q2)
}
