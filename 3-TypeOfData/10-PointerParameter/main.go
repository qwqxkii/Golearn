package main

import "fmt"

func main() {
	var num int = 100
	Double(num)
	fmt.Println(num)

	Doublewithptr(&num)
	fmt.Println(num)

	println("分割")

	var a1 = [3]int{100, 200, 300}
	Aarraryptr(a1)
	fmt.Println(a1)

	Arraryptr2(&a1)
	fmt.Println(a1)

	println("分割")
	var s1 = []int{1, 2, 3}
	Sliceptr(s1)
	fmt.Println(s1)
}

// 两个func都没有返回哈，也就是说第一个func内部的n*=n 应该是200但只会影响func内部
func Double(n int) {
	n *= n
}

// 这个才会影响全局的变量
func Doublewithptr(n *int) {
	*n *= 2
}

func Aarraryptr(n [3]int) {
	for i, _ := range n {
		n[i] *= 2
		// println(n[i])   200 400 600
	}
}
func Arraryptr2(n *[3]int) {
	for i, _ := range n {
		n[i] *= 2
	}
}

// 切片参数  //切片相当于一个数组  //指向了底层数组的元素
func Sliceptr(n []int) {
	for i, _ := range n {
		n[i] *= 2
	}
}
