package main

import "fmt"

func main() {
	//for初始化条件；判断条件；条件编号
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
