package main

import (
	"fmt"
	"sort"
)

func main() {
	ordermap()
}
func ordermap() {
	//创建一个map
	var m1 = make(map[int]string, 4)
	//给map赋值
	m1[1] = "one"
	m1[0] = "zero"
	m1[2] = "two"
	m1[3] = "three"

	fmt.Println(m1)
	//创建一个slice用于存放map的key用于排序
	s1 := make([]int, len(m1))
	//创建一个索引
	index := 0
	for i, _ := range m1 {
		s1[index] = i
		index++
	}
	fmt.Println(s1)

	//排序
	sort.Ints(s1)
	//打印出来
	fmt.Println(s1)

	//循环五次用于验证
	for q := 0; q < 4; q++ {
		//当然要循环输出
		for _, value := range s1 {
			fmt.Printf("id :%d, value:%s\n", value, m1[value])
		}
		fmt.Println("第", q+1)
	}
}
