package main

import "fmt"

// 数组
func main() {
	//数组
	var a1 = [3]int{} //3个类型为int的数组。默认值为0
	//切片
	var s1 = []int{}
	fmt.Println(a1, s1)
	demo1()
	a1[1] = 3
	fmt.Println(a1)
	demo2()
	forarray()

	demo99()
}
func demo1() {
	var str = string("1q2q3")
	fmt.Printf("%T,%v\n", str, str)

}
func demo2() {
	//不定数组
	var a1 = [...]int{0: 9, 13: 99}
	a1[0] = 2
	a1[13] = 199
	fmt.Println(a1, len(a1))
}

// 遍历数组
func forarray() {
	var s1 = [6]int{3, 2, 3, 5, 2}
	fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d,%d\n", i, s1[i])
	}
	//range
	for i, v := range s1 {
		fmt.Printf("%dm%d\n", i, v)
	}
}

// 数组打印99乘法表
func demo99() {
	var arr [9][9]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = i*3 + j + 1
		}
	}

	//输出
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d * %d = %2d\t", i+1, j+1, arr[i][j])
		}
		fmt.Println()
	}
}
