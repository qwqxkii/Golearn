package main

import "fmt"

type persion struct {
	name     string
	age      int
	address  string
	annimols annimols
	memony   memony
}
type annimols struct {
	age  int
	name string
}
type memony struct {
	dollars int
}

func main() {
	var p1 persion
	p1.name = "小三"
	p1.age = 19
	p1.address = "老北京"
	p1.annimols.name = "小黑"
	p1.annimols.age = 3
	p1.memony.dollars = 200
	fmt.Println(p1)

	var p2 = persion{name: "西奥斯", age: 23, address: "上海", annimols: annimols{name: "大黑", age: 2}, memony: memony{dollars: 200}}
	fmt.Println(p2)
}
