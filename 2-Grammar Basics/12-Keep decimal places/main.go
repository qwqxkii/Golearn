package main

import (
	"fmt"
	"math"
)

func main() {
	var f1 float64 = 1.32232323
	s1 := fmt.Sprintf("%.3f", f1) //保存3位
	fmt.Printf("%T,%v\n", s1, s1)

	s2 := fmt.Sprintf("%.1f", f1)
	fmt.Printf("%T, %v\n", s2, s2) //保存1位

	demo()
}
func demo() {
	f1 := 12.3323
	var rain float64
	rain = math.Pow(10, 2) //保留两位
	s1 := math.Round(f1*rain) / rain
	fmt.Printf("%T,%v\n", s1, s1)

	rain = math.Pow(10, 1) //保留1位
	s1 = math.Round(f1*rain) / rain
	fmt.Printf("%T,%v\n", s1, s1)

}
