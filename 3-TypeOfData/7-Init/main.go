package main

//Import->const->var->init()->main()
import "fmt"
import "learn/3-TypeOfData/7-Init/helpers"

var (
	zhang int
	fei   int
)

func init() {
	zhang = 2
	fei = 99
}

func main() {
	fmt.Println(fei, zhang)
	sum := helpers.Add(5, 3)
	sub := helpers.Sub(5, 3)
	fmt.Println(sum, sub)

}
