package main

import "fmt"

func main() {
	x, xx := func(a, b int) (max, min int) {
		if a > b {
			max = a
			min = b
		} else {
			max = b
			min = a
		}
		return
	}(10, 20)
	fmt.Println(x, xx)
	fmt.Println("testxxx", "jud")
}
