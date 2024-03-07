package main

import (
	"fmt"
	"sort"
)

func main() {
	demo1()

}
func demo1() {
	//创建
	var m1 = make(map[int]string, 5)

	//增加值
	m1[1] = "张山"
	m1[2] = "夏鹏"
	m1[5] = "李金"
	m1[3] = "物资"

	//创建一个切片
	str := make([]int, len(m1))
	index := 0

	//按照索引放进去
	for i, _ := range m1 {
		str[index] = i
		fmt.Println(str[index], i)
		index++
	}
	//排序
	sort.Ints(str)
	//分隔
	fmt.Println(str, "2222222")
	//循环5次，用于验证
	for i := 0; i < 5; i++ {
		//循环读出切片
		for _, key := range str {
			fmt.Printf("key:%d,value:%s\n", str, m1[key])
		}
		fmt.Printf("第 %d 次遍历完成\n", i+1)
	}
}
