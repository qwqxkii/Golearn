// 练习题 1：实现接口

package main

import "fmt"

//定义一个 Vehicle 接口，它包含一个方法 Drive()，该方法不返回任何值。
type Vehicle interface {
	Drive()
}

//然后定义两个结构体：Car 和 Bike。每个结构体都应实现 Vehicle 接口的 Drive 方法。
type Car struct {
}

// 在 Drive 方法中，让 Car 输出 "Driving a car"，而 Bike 输出 "Riding a bike"。
func (c *Car) Drive() {
	fmt.Println("Driving a car")
}

type Bike struct {
}

func (c *Bike) Drive() {
	fmt.Println("Riding a bike")
}

func main() {
	t1 := new(Car)
	t2 := new(Bike)
	t1.Drive()
	t2.Drive()
}
