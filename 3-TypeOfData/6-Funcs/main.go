package main

import "fmt"

func main() {
	f1 := demo1(1, 2)
	fmt.Println(f1)
	f2, f3 := demo2()
	fmt.Println(f2, f3)
	f4, f5 := demo3()
	fmt.Println(f4, f5)

	f6, f7 := demo4(100, 200, 300)
	fmt.Println(f6, f7)

}
func funcdemo1(v int, q int) int {
	return v + q
}

func funcdemo2(v int, q int) (int, int) {
	return v, q
}

func funcdemo3(v int, q int) (int, int, int) {
	return v, q, v + q
}

func demo1(x, y int) int {
	return x + y
}
func demo2() (int, int) {
	return 1, 2
}
func demo3() (x int, y int) {
	x, y = 30, 2
	return x, y
}
func demo4(x, y, z int) (max, min int) {
	max = x
	min = x

	if y > max {
		max = y
	} else if y < min {
		min = y
	}

	if z > max {
		max = z
	} else if z < min {
		min = z
	}
	return max, min
}
