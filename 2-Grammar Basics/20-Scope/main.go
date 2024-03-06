package main

// 作用域

var q int = 100 //全局变量

func main() {

	//局部变量
	for i := 0; i < 3; i++ {
		println(i, "i只能在这个循环里使用。")
	}

	//println(i)  //为解析的引用

	println(q, "引用全局变量")

	//var zhang int = 20
	for i := 0; i < 3; i++ {
		//优先引用内部变量
		var zhang int = 1
		println(zhang, "i只能在这个循环里使用。")
	}

}
