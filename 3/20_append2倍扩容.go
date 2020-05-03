package main

import "fmt"

func main() {
	//如果超过原来的容量，通常以2倍容量扩容
	s := make([]int, 0, 3)
	oldCap := cap(s)

	for i := 0; i < 100; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap != newCap {
			fmt.Printf("i:%d, newcap:%d, oldcap:%d\n", i, newCap, oldCap)
			oldCap = newCap
		}
	}
}
