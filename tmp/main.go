package main

import "fmt"

func main() {
	//闭包
	f1 := Closure("小九")
	fmt.Println(f1())
}
func Closure(name string) func() string {
	return func() string {
		return name
	}
}
