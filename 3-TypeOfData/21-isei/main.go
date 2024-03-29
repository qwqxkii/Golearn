package main

import "fmt"

// 实现系统错误接口
func main() {
	divde, err := myDivde(100, 2)
	if err != nil {
		fmt.Printf("ERROR:%s\n", err)
	} else {
		fmt.Printf("100/0,%2.f\n", divde)
	}
}

// 自定义核心除法函数
// 最后完善核心函数
func myDivde(divdend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, newDivideError()
	}
	return divdend / divisor, nil
}

// 定义错误类型
type divideError struct {
	msg string
}

func (d *divideError) Error() string {
	return d.msg
}

// 创建自定义错误实例
func newDivideError() *divideError {
	return &divideError{msg: "有毛病吧"}
}
