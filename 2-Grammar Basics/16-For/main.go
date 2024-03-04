package main

func main() {
	demo1()
	println("======")
	demo2()
	println("======")
	demo3()
}
func demo1() {
	for i := 0; i < 3; i++ {
		println(i)
	}
}
func demo2() {
	i := 0
	for i < 3 {
		i++
		println(i)
	}
}
func demo3() {
	i := 0
	for i < 5 {
		i++
		if i > 3 {
			break
		}
		println(i)
	}
}
func demo4() {
	//无限循环
	//for ()
}
