package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	//fmt.Print(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	//fmt.Println()
}
