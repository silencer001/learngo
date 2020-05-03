package main

import "fmt"

func test(i int) (sum int) {
	if i == 1 {
		return 1
	}
	return i + test(i-1)
}
func main() {
	fmt.Println(test(100))
}
