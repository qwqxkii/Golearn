package main

import "fmt"

func main() {
	f1 := Recursion(5)
	fmt.Println(f1)
}
func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n + Recursion(n-1)
}
