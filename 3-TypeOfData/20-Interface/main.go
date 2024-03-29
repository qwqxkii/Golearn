package main

import (
	"errors"
	"fmt"
)

func main() {
	one := &Cat{name: "cat"}
	println(one.Speak())

	println("------")
	two := &Box{b1: 2, b2: 4}
	fmt.Printf("%2.f \n", two.Area())

	println("------")
	three := &Pencil{}
	err := three.Write("test")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	//手动报错
	err = three.Write("error")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

}

// 创建一个结构体实现 Animal 接口，实现 Speak() 方法并返回字符串 "Meow"，然后调用 Speak() 方法并输出结果。
type Animal interface {
	Speak() string
}
type Cat struct {
	name string
}

func (a *Cat) Speak() string {
	return a.name + "Meow"
}

// 创建一个结构体实现 Shape 接口，实现 Area() 方法计算一个正方形的面积，并返回结果，然后调用 Area() 方法并输出结果。
type Shape interface {
	Area() float64
}
type Box struct {
	b1 float64
	b2 float64
}

func (b *Box) Area() float64 {
	return b.b1 * b.b2
}

// 创建一个结构体实现 Writer 接口，实现 Write() 方法将传入的字符串打印到控制台，并返回 nil，然后调用 Write() 方法并检查是否有错误。
type Writer interface {
	Write(data string) error
}
type Pencil struct {
}

func (p *Pencil) Write(data string) error {
	if data == "error" {
		return errors.New("Find a new ERROR!")
	}
	return nil
}
