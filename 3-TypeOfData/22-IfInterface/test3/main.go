package main

import "fmt"

//定义一个接口 Painter，它包含一个方法 Paint() 返回一个字符串表示绘制的内容。
type Painter interface {
	Paint() string
}

//定义两个结构体 Artist 和 Child，2.它们都实现了 Painter 接口。
type Artist struct {
	name string
}
type Child struct {
	name string
}

// Artist 的 Paint 方法返回 "Painting a masterpiece"，而 Child 的 Paint 方法返回 "Painting for fun"。
func (a *Artist) Paint() string {
	return a.name + "Painting a masterpiece"
}

func (a *Child) Paint() string {
	return a.name + "Painting for fun"
}

//编写一个函数 displayArtwork，它接受 Painter 类型的参数并调用其 Paint 方法，然后输出返回的字符串。
func displayArtwork(p Painter) string {
	return p.Paint()
}

func main() {
	var i1 = Artist{name: "小王"}
	var i2 = Child{name: "肖利茨"}

	fmt.Println(displayArtwork(&i1))
	fmt.Println(displayArtwork(&i2))
}

// 练习题 3：多态和接口
