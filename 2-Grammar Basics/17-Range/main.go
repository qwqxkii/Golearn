package main

import "fmt"

func main() {
	demo()
	demo2()
}
func demo() {
	var str1 = "123xx2332"
	//for _, v := range str1 {  隐藏index
	//for i, _ := range str1 {  隐藏value
	for i, v := range str1 {
		fmt.Printf("index:%d value:%c\n", i, v)
	}
}
func demo2() {
	str := "hello"
	for _, _ = range str {
		println("do something")
	}

}
