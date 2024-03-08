package main

import (
	"fmt"
	"sort"
)

func main() {
	demo()
}
func demo() {
	//创建一个map
	var m1 = make(map[int]string)
	//给map的key和value赋值
	m1[4] = "张飞"
	m1[2] = "老王"
	m1[3] = "网布"
	m1[1] = "张牛"
	//创建一个空int切片
	s1 := make([]int, len(m1))
	//创建一个索引
	index := 0
	//循环将map的key赋值到切片中
	for i, _ := range m1 {
		s1[index] = i //把map扔进[]int切片中
		index++
	}
	//输出看看
	fmt.Println(s1, "---")
	//切片排序&输出看看
	sort.Ints(s1)
	fmt.Println(s1)
	//循环输出5次-验证
	for i := 0; i <= 4; i++ {
		for _, qwq := range s1 {
			fmt.Printf("key %d value %s\n", qwq, m1[qwq]) //自然用切片输出

		}
		println("-------")
	}

}
