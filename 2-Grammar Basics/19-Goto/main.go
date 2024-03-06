package main

import "os"

// goto
func main() {
	demo1()
}
func demo1() { //首字母大写，可导出，首字母小写，不可导出。这个demo1开头小写只能在本包内访问
	a := 0
LABEL:
	println("看我脸色行事")
	a++
	if a > 3 {
		os.Exit(0)
	}
	goto LABEL //跳到对应标签位置。
}

// 公共包
func Demo2() {

}
