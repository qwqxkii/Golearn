package main

import (
	"errors"
	"fmt"
)

// 练习题 1: 使用 errors.New()
// 题目：你正在编写一个函数CheckUsername，该函数接收一个用户名（字符串）作为输入。如果用户名小于5个字符，函数应返回一个错误，表明用户名太短。请使用errors.New()创建并返回这个错误。
func CheckUsername(name string) (int, string, error) {
	if len(name) < 5 {
		return len(name), "用户名检测未通过", errors.New("用户名少于5个字符")
	}
	return len(name), "用户名检测通过", nil
}
func main() {
	n1 := "小王子"
	n2 := "xiaowangzi"
	str, num, err := CheckUsername(n1)
	fmt.Println(str, num, err)

	str, num, err = CheckUsername(n2)
	fmt.Println(str, num, err)
}
