package main

import "fmt"

type divError struct {
	msg string
}

func (d *divError) Error() string {
	return d.msg
}

func newError() *divError {
	return &divError{msg: "这是自定义错误哈"}
}
func Math(a, b int) (c int, err error) {
	if b == 0 {
		return 0, newError()
	}
	return a + b, nil
}

func main() {
	str, err := Math(1, 5)
	if err != nil {
		fmt.Println("出现问题请查看:", str, err)
	} else {
		fmt.Println("未发现错误，正常输出：", str)
	}

}
