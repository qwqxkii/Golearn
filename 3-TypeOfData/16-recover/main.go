package main

import "fmt"

func main() {
	test()
}
func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("发现一个错误%v", r)
		}
	}()
	panic("error")
}
