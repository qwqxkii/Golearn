package main

import "fmt"

type AandB interface {
	add(float64) float64
	del(float64) float64
}
type people struct {
	name string
}
type pets struct {
	name string
}

func (p *people) add(num float64) float64 {
	return num + 100
}
func (p *people) del(num float64) float64 {
	return num - 10
}

func main() {
	//判断people是否满足AandB接口
	var i1 interface{}
	i1 = &people{name: "校长"}
	if v, ok := i1.(AandB); ok {
		fmt.Println("满足", v.add(1), v.del(22))
	} else {
		fmt.Println("不满足接口")
	}
	var i2 interface{}
	i2 = &pets{
		name: "小儿",
	}
	if v, ok := i2.(AandB); ok {
		fmt.Println("满足接口", v.del(22))
	} else {
		fmt.Println("不满足接口")
	}
}
