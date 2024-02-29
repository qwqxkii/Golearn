package main

//这个fmt就是被导入的包
import (
	fmt "fmt"
	//xx "fmt"
	"strings"
)

// 导入包 #
func main() {
	//fmt包里面有个Println
	fmt.Println("fuvk you ass")
	//
	//xx.Println("fuvk you ass")
	//strings.Repeat 是 Go 语言中 strings 包提供的一个函数，用于重复指定字符串若干次并返回结果。函数的签名为
	//func Repeat(s string, count int) string，其中 s 是要重复的字符串，count 是重复的次数。
	fmt.Println(strings.Repeat("s", 3)) //输出sss
}
