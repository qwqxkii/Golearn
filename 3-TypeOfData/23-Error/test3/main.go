package main

import (
	"fmt"
)

//练习题 3: 使用 fmt.Errorf()

// 题目：你需要完成一个名为Divide的函数，该函数接收两个整数a和b，并返回它们的除法结果。
func Divide(a, b int) (int, error) {
	//如果b为0，函数应返回一个错误，说明除数不能为0。
	if b == 0 {
		////请使用fmt.Errorf()来创建并返回这个错误，同时在错误信息中包含b的值。
		return 0, fmt.Errorf("值不能为0,目前你传递的值是:%d", b)
	}
	return a / b, nil
}
func main() {
	value, err := Divide(10, 0)
	if err != nil {
		fmt.Println("有问题", value, err)
	} else {
		fmt.Println("计算结果", value)
	}
}
