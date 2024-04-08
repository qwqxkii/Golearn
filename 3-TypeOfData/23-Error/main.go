package main

import (
	"errors"
	"fmt"
)

// 第一种
func divErro(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("不能为零")
	}
	return a / b, nil
}

// 第二种
func divErrorf(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("你传入的值不能为0 %.2f", b)
	}
	return a / b, nil
}
func main() {
	div, err := divErro(10, 0)
	if err != nil {
		fmt.Println("返回的值:", div, "有错误:", err)
	} else {
		fmt.Println("没问题,返回的值", div)
	}
	div, err = divErrorf(10, 0)
	if err != nil {
		fmt.Println("返回的值:", div, "有错误:", err)
	} else {
		fmt.Println("没问题,返回的值", div)
	}
}

//errors.New() 创建错误 #
