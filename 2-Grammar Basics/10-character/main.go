package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := 'a'                                              //"这是字符串"        '这是字符'
	fmt.Printf("s type:%T, len:%d\n", s, utf8.RuneLen(s)) //len(s)无法接收

	s2 := '俺'
	fmt.Printf("s2 type:%T len:%d\n", s2, utf8.RuneLen(s2))
}
