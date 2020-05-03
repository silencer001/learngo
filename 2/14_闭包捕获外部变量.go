package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	//闭包以引用方式捕获外部变量，闭包函数内更改，会影响到外部变量
	func() {
		a = 666
		str = "go"
		fmt.Printf("内部:a=%d, str=%s\n", a, str)

	}()
	fmt.Printf("外部：a=%d, str=%s\n", a, str)
}
