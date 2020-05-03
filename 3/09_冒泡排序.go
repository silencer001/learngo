package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bable(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	//冒泡排序
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
