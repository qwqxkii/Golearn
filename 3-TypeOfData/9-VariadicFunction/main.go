package main

import "fmt"

func main() {
	q := Add(1, 2, 3, 4, 5, 5, 2, 2)
	q2 := Add(1)
	fmt.Println(q, q2)

	s1 := []int{1, 2, 3}
	s2 := SliceAdd(s1...)
	fmt.Printf("type: %T %T value: %d\n", s2, s1, s2)

	//不存传递任何参数
	fmt.Println(SliceAdd())
}
func Add(q ...int) int {
	sum := 0
	for _, k := range q {
		sum += k
	}
	return sum
}
func SliceAdd(q ...int) int {
	sum := 0
	for _, k := range q {
		sum += k
	}
	return sum
}
