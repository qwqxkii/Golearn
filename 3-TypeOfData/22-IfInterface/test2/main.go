// 练习题 2：接口类型断言
package main

import "fmt"

//创建一个接口 Movable，包含一个方法 Move()。
type Movable interface {
	Move()
}

//定义两个结构体 Robot 和 Human，并实现 Movable 接口。
type Robot struct {
}

func (r *Robot) Move() {
	fmt.Println("Robot is moving")
}

type Human struct {
}

func (h *Human) Move() {
	fmt.Println("Human is moving")
}

// 编写一个函数 identifyMovableType，它接受 Movable 类型的参数并使用类型断言来判断传入的参数是 Robot 还是 Human 类型。根据类型打印出不同的消息。
func identifyMovableType(m Movable) {
	if _, ok := m.(*Robot); ok {
		fmt.Printf("这是robot类型\n")
	} else if _, ok := m.(*Human); ok {

		fmt.Printf("这是human类型\n")

	} else {
		fmt.Println("异常！！！")
	}
}
func main() {
	var i2 = new(Robot)
	var i1 = new(Human)
	i1.Move()
	i2.Move()
	identifyMovableType(i2)
	identifyMovableType(i1)

}
