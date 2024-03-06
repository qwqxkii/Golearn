package main

import "fmt"

func main() {
	Mapdemo()
	println()
	DeleteMap()
}
func Mapdemo() {
	m1 := make(map[int]int, 3) //使用make申请内存
	m1[1] = 2
	m1[2] = 9
	fmt.Println(m1, len(m1))
}
func DeleteMap() {
	m1 := make(map[int]string, 3)
	m1[1] = "张山"
	m1[2] = "里斯"
	m1[3] = "向往"
	fmt.Println(m1, len(m1))
	println()
	delete(m1, 1)

	fmt.Println(m1, len(m1))

}
