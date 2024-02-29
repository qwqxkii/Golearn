package main

import "fmt"

func main() {
	//常量的值必须在编译时就能确定
	const q1 = 1
	const q2 = 2
	const q3 = 3
	const q4 = 4
	const q5 = 5 //常量初始化后不使用也不会报错

	fmt.Println(q3)
}
