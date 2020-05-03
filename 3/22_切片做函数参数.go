package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bable(a []int) {
	for i := 1; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
	fmt.Println("InitData :", s)
}
func main() {
	var n int = 10
	//创建一个切片，长度为n
	s := make([]int, n)

	InitData(s) //初始化slice
	fmt.Println("排序前:", s)
	bable(s)
	fmt.Println("排序后：", s)
}
