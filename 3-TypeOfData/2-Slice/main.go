package main

import "fmt"

func main() {
	var s1 []int

	sm1 := make([]int, 3)
	sm1[2] = 3
	fmt.Printf("%T ,%v,长度%d 容量%d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("%T ,%v,长度%d 容量%d\n", sm1, sm1, len(sm1), cap(sm1))
	SliceAndArray()
	AppendSlice()
	Forarray()
}

// 多个切片引用一个数组
func SliceAndArray() {
	var v1 = [6]int{3, 2, 12, 1, 3, 45}
	s1 := v1[3:5]
	s2 := v1[:3]
	s3 := v1[4:]
	s4 := v1[1:5]
	fmt.Println(v1, s1, s2, s3, s4)
	v1[0] = 99 //修改前是3
	fmt.Println(v1)

}

// 追加值
func AppendSlice() {
	var s1 []int
	fmt.Printf("%d,%d\n", len(s1), cap(s1))
	s1 = append(s1, 3)
	fmt.Printf("%d,%d\n", len(s1), cap(s1))
}

// 遍历切片
func Forarray() {
	var v1 = [6]int{3, 2, 12, 1, 3, 45}
	for i, v := range v1 {
		fmt.Printf("索引:%d值:%d\n", i, v)
	}
}
