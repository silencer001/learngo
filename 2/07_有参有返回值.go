package main

import "fmt"

func maxandmin(a, b int) (max int, min int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	fmt.Println(maxandmin(1, 3))
}
