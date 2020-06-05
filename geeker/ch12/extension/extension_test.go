package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//p *Pet
	Pet //匿名方法类型
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Print("Wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	//d.p.Speak()
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {
	var dog *Dog = new(Dog)
	dog.SpeakTo("ju")
}
