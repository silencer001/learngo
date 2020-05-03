package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a  = %d, b = %d \n", a, b)
}
