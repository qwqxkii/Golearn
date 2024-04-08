package main

import "fmt"

type Speaker interface {
	say(string)
}
type Persion struct {
	name string
	age  int
}

func SpeakerMan(speaker Speaker) {
	speaker.say("wwwww")
}
func (p *Persion) say(msg string) {
	fmt.Println(p.name, msg)
}

func main() {
	p1 := new(Persion)
	p1.name = "小王"
	p1.age = 18
	fmt.Println(p1)
	p1.say("好")

	SpeakerMan(p1)

}
