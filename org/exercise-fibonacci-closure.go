package main

import "fmt"

func fibonacci() func() int {
	n := 0
	fn := 0
	fn_1 := 0
	fn_2 := 0
	return func() int {
		switch n {
		case 0:
			fn = 0
		case 1:
			fn = 1
		default:
			fn_2 = fn_1
			fn_1 = fn
			fn = fn_1 + fn_2
		}
		n++
		return fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
