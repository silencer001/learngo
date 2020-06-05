package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestGoProgrammer(t *testing.T) {
	var i Programmer
	i = new(GoProgrammer)
	t.Log(i.WriteHelloWorld())
}
