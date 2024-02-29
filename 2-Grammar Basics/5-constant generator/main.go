package main

// constant generator常量生成器 #
func main() {
	constIota()
}

func constIota() {
	const (
		q1 = iota
		q2
		q3
		q4
		q5
	)

	println(q1, q2, q3, q4, q5)
}
