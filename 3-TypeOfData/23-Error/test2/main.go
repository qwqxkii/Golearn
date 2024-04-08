package main

import (
	"errors"
	"fmt"
	"os"
)

// 练习题 2: 结合使用 errors.New() 和错误检查
// 题目：编写一个OpenFile函数，该函数接收文件名（字符串）作为参数。
func OpenFile(filename string) error {
	//使用os.Open尝试打开指定的文件。
	f, err := os.Open(filename)
	//如果os.Open返回错误，使用errors.New()创建一个新的错误
	if err != nil {
		//，表明文件无法打开，并返回这个错误。
		return errors.New("文件无法打开")
	}
	defer f.Close()
	return nil
}
func main() {
	err := OpenFile("qqwq.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件操作成功。")
	}
}
