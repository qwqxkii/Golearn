package main

import "fmt"

func main() {
	fmt.Println("环境测试")
	//&取地址符    *根据地址取值
	var qwq int = 20
	var fuvk string = "张飞"
	fmt.Printf("%T,%T,%v,%v \n", qwq, fuvk, qwq, fuvk)
	q1 := &qwq
	f1 := &fuvk
	//根据&取地址，取到内存的地址，相当于门牌号
	fmt.Printf("%p,%p\n", q1, f1)

	//根据地址取值
	q2 := *q1
	fmt.Printf("%T,%v", q2, q2)

}
