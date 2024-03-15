package main

import "fmt"

func bibao() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	bi := bibao()
	fmt.Println(bi())
	for i := 1; i < 5; i++ {
		fmt.Println(i, bi())
	}

}
