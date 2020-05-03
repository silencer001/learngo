package main

import "fmt"

func main() {
	//iota常量自动生成器，每个一行，自动累加1
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a= %d, d= %d, c=%d\n", a, b, c)
	//iota遇到cosnt关键字，重置为0
	const d = iota
	fmt.Printf("d = %d\n", d)
	//const e := iota
	//fmt.Printf("e=%d\n", e)
	//结论：iota不能赋值给变量

	// 可以只写一个iota
	const (
		a1 = iota
		b1 = iota
		c1
	)
	fmt.Printf("a1=%d, b1=%d, c1=%d\n", a1, b1, c1)

	//如果是同一行，赋值一样
	const (
		a2 = iota
		f2
		b2, c2, d2 = iota, iota, iota
		e2         = iota
		g2
	)
	fmt.Printf("a2=%d, b2=%d, c2=%d\n, d2=%d, e2=%d\n", a2, b2, c2, d2, e2)
}
