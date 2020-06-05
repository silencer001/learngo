package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.Println(\"Hello World\")"
}
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
func TestGoProgrammer(t *testing.T) {
	var i Programmer
	i = new(GoProgrammer)
	t.Log(i.WriteHelloWorld())
	i = new(JavaProgrammer)
	t.Log(i.WriteHelloWorld())

	goer := new(GoProgrammer)
	javaer := new(JavaProgrammer)
	writeFirstProgram(goer)
	writeFirstProgram(javaer)
}
