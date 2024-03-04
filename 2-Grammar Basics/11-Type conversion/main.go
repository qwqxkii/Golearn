package main

import (
	"fmt"
	"strconv"
)

// 类型转换
// go不支持隐式类型转换
func main() {
	ChangeTypeDemo1()
	ChangeTypeDemo2()
	ChangeTypeDemo3()

	demo1()
}
func ChangeTypeDemo1() {
	n := 100
	floatN := float64(n)
	fmt.Println(float64(floatN))
}
func ChangeTypeDemo2() {
	n := 100
	fmt.Println(float64(n))
}
func ChangeTypeDemo3() {
	n := 3.1
	fmt.Println(int64(n))

}

func demo1() {
	s := strconv.Itoa(1024)
	fmt.Printf("%T, %v\n", s, s) //整形转为字符串

	n, _ := strconv.Atoi("1024")
	fmt.Printf("%T, %v\n", n, n) //字符串转为整行

	s2 := strconv.FormatFloat(3.15323, 'f', -1, 64) //浮点转为字符串
	fmt.Printf("%T, %v\n", s2, s2)

	n2, _ := strconv.ParseFloat("3.12323", 64)
	fmt.Printf("%T, %v\n", n2, n2) //字符串转为浮点类型
}
