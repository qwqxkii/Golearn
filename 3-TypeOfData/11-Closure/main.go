package main

import "fmt"

func inSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	next := inSeq()
	fmt.Println(next())
	for i := 0; i < 3; i++ {
		fmt.Println(next(), i)
	}
}
