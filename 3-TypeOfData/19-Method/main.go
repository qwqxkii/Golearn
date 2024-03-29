package main

import "fmt"

// 方法

type persion struct {
	name string
	age  int
}

func main() {
	Method()
	Demo2()
}
func Method() {
	p1 := &persion{name: "小子", age: 12}
	p1.Sayhello()

	p1.SetName("老王")
	fmt.Println(p1)
}

func (p persion) Sayhello() {
	fmt.Println(p.name + "你好宰总")
}
func (p *persion) SetName(newname string) {
	p.name = newname
	fmt.Println("老名字" + p.name + "新名字" + newname)
}

func Demo2() {
	println()
	p1 := &persion{name: "小子里", age: 11}
	fmt.Println(p1)
	p1.Run()
	p1.SetName("老丸子")
	p1.Sayhello()
}
func (p *persion) Run() {
	fmt.Println(p.name + "在跑")
}
func (p *persion) Walk() {
	fmt.Println(p.name + "在走")
}
