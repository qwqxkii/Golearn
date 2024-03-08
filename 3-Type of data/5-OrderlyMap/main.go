package main

import (
	"fmt"
	"sort"
)

func main() {
	MapOrder()
}
func MapOrder() {
	var m1 = make(map[int]string)

	m1[2] = "小屋"
	m1[5] = "小三"
	m1[1] = "异性"
	m1[3] = "学习"
	m1[4] = "高级"

	var s1 = make([]int, len(m1))
	var index = int(0)

	for i, _ := range m1 {
		s1[index] = i
		index++
	}

	sort.Ints(s1)
	fmt.Println(s1)

	for i := 1; i < len(s1); i++ {
		for _, q := range s1 {
			fmt.Printf("key %d value %s\n", q, m1[q])
		}
		fmt.Println(i + 1)
	}
}
