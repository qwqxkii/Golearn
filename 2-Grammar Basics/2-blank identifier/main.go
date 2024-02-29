package main

import "strconv"

func main() {
	var qwq = 2
	_ = qwq //引用然后空调用

	//样例
	n, _ := strconv.Atoi("1024") //正常来说这个trconv.Atoi会返回两个数值，转换后的整数会被打印到控制台
	//第二个返回是出现错误，将打印错误信息,在本样例中，忽略接收第二个
	println(n)
}
