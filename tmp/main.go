package main

import "fmt"

func MaxMin(a, b, c int) (max, min int) {
	max = a
	min = a
	if b > max {
		max = b
	} else if b < max {
		min = b
	}
	if c > max {
		max = c
	} else if c < max {
		min = c
	}
	return max, min
}
func main() {
	f1, f2 := MaxMin(100, 200, 300)
	fmt.Printf("最大值%d，最小值%d\n", f1, f2)
}
