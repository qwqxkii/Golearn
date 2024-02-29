package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "cheese\nsalad" //换行
	println(str1)
	//\n 换行 \t tab \r 回车 \\反斜  \u unicode

	str2 := `听说会原样输出
，那这不得试试？`
	println(str2)

	str3 := "welcome the fuvk world"
	//字符串不可变,会报错
	//str3[0] = 'a'
	//str3[1] = 'a' //无法分配给str3[1]      ./main.go:6:2: cannot assign to s[1] (value of type byte)
	println(str3)

	//可以重新赋值
	str4 := "welcome the fuvk world"
	str4 = "fuvk me very hard"
	println(str4)

	println(len(str4)) //len(s) 返回字符串变量 s 中字符的个数

	str5 := "中文算三个"
	println(len(str5))

	println("=====")
	//那如何修改为中文算一个字符统计
	//import "unicode/utf8"
	str6 := "hello world"
	println(utf8.RuneCountInString(str6))

	str7 := "hello 世界"
	println(utf8.RuneCountInString(str7))

	substring()

	traverseAstring()

	stringConcatenation()
}
func substring() {
	s := "hello world"
	s1 := s[2:8]
	fmt.Printf("%v,类型%T\n", s1, s1)

	s2 := s[:]
	fmt.Printf("%v,%T\n", s2, s2)
}
func traverseAstring() {
	str := "qqq1122qqq"
	for i, v := range str {
		fmt.Printf("index:%d, value:%c\n", i, v)
	}
	println("==========")
	for i := 0; i < len(str); i++ {
		fmt.Printf("索引%d,值%c \n", i, str[i])
	}
}
func stringConcatenation() {
	str1 := "这是"
	str1 += "?"
	println(str1)

	//使用 fmt.Sprintf() 函数
	str2 := fmt.Sprintf("%s %s", str1, "!")
	println(str2)
}
