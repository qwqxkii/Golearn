package main

import "fmt"

// 定义一个名为Shaper的接口，它包含一个名为Area()的方法，该方法返回float64类型的值。
type Shaper interface {
	Area() float64
}
type God struct {
}

func (g *God) Area() float64 {
	return float64(222)
}

func main() {
	g := &God{}
	a := g.Area()
	fmt.Println(a)
}
