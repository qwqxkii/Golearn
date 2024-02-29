package main

import "fmt"

func main() {
	//普通打印 #
	//优点：内置函数，不需要引入额外的包，简单方便。
	//不足：无法进行格式化打印，无法完整打印复合数据结构 (如数组, Map 等)。
	println(1024, "hello world", true)

	//格式化打印 #
	//这里先介绍 2 个方法，分别是 fmt 包里面的 Println() 和 Printf(), 大多数场景下都适用。
	fmt.Println(1024, "hello jimi", true)

	//最重要的格式化打印函数之一来了
	//https://golang.dbwu.tech/introduction/print/
	//fmt.Printf()
	fmt_demo()
	println("===")
	fmt_demo2()
}
func fmt_demo() {
	n := 1024
	fmt.Printf("这是整数%d\n", n)

	pi := 3.1415926897
	fmt.Printf("这是浮点数%f\n", pi)

	str := "抱你狗头"
	fmt.Printf("这是字符串%s\n", str)

	yes := true
	fmt.Printf("t这是布尔%t\n", yes)

	x := 17
	fmt.Printf("这是二进制%b", x)
}
func fmt_demo2() {
	n := 1024
	fmt.Printf("%T %d %v\n", n, n, n) //看最后，是不是要输入3个n
	//优化
	fmt.Printf("%T %[1]d %[1]v\n", n) //这是优化
}
