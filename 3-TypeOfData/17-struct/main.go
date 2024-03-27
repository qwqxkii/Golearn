package main

import "fmt"

type persion struct {
	name string
	age  int
}

func main() {
	var p1 persion
	p1.name = "这三个"
	p1.age = 13
	fmt.Printf("%v,%T", p1, p1)
	println()
	fmt.Println(p1.name, p1.age)

	p2 := persion{}
	p2.name = "111"
	p2.age = 11
	fmt.Println(p2)

	p3 := new(persion)
	p3.name = "区别？"
	p3.age = 13
	fmt.Printf("%v,%T", p3.name, p3)

}
