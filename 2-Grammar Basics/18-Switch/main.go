package main

func main() {
	a := -0
	switch {
	case a >= 100 && a < 200:
		println("111")
	case a > 200:
		println("222")
	default:
		println("error")
	}
	demo1()
}
func demo1() {
	a := 10
	switch a {
	case 10:
		println("11")
	case 22:
		println("22")
	default:
		println("??")
	}
}
