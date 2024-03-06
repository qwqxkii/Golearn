package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	CharSlice()
	byteTostring()
	LongCount()
	ChinaCount()
	ChinaCountOne()
}
func CharSlice() {
	str1 := "hello world"
	s := []byte(str1)
	fmt.Printf("%T,%[1]s\n", s)
}
func byteTostring() {
	var b1 = []byte{'h', '3', '4', 'h'}
	str := string(b1)
	fmt.Printf("%T,%s\n", str, str)
}

// 长度计算 #ASCII
func LongCount() {
	var b1 = []byte{'1', 'x', '@', '2'}
	fmt.Printf("b1 %s of lang %d\n", b1, len(b1))

}

// 中文算作 3 个字符 #
func ChinaCount() {
	b1 := []byte("俺是")
	fmt.Printf("b length = %d\n", len(b1))
}
func ChinaCountOne() {
	b1 := []byte("俺是")
	fmt.Printf("b length = %d\n", utf8.RuneCount(b1))
}
