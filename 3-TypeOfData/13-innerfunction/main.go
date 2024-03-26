package main

import "fmt"

// 内部求和函数
func main() {
	var sum func(...int) int //定义函数

	sum = func(manynum ...int) int {
		total := 0
		for _, v := range manynum {
			total += v
		}
		return total
	}
	fmt.Println(sum(1, 2, 3, 4, 5))
}
