package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:4:5]
	fmt.Printf("s: %v", s)
	fmt.Printf("len(s)= %d, cap(s)= %d\n", len(s), cap(s))
}
