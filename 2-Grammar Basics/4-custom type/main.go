package main

type fzh int

// 再嵌套
type xue fzh

func main() {
	var qwq fzh = 23
	println(qwq)

	var xzx xue = 22
	println(xzx)
}
