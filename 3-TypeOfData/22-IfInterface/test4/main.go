package main

import "fmt"

//定义两个接口 Writer 和 Reader，Writer 包含 Write() 方法，Reader 包含 Read() 方法。
type Writer interface {
	Write()
}

type Reader interface {
	Read()
}

//然后定义一个新的接口 Editor，它组合了 Writer 和 Reader 接口。
type Editor interface {
	Writer
	Reader
}

//定义一个结构体 TextEditor 并实现 Editor 接口。
type TextEditor struct {
}

//2.实现
func (t *TextEditor) Write() {
	fmt.Println("Writing text")
}
func (t *TextEditor) Read() {
	fmt.Println("Reading text")
}

func main() {
	//使用
	var str1 = TextEditor{}
	str1.Read()
	str1.Write()
}

// 练习题 4：接口组合
// TextEditor 的 Write 方法输出 "Writing text"，Read 方法输出 "Reading text"。
