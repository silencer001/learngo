package main
import "fmt"
var a int = 10
//全局变量不能使用简单定义
// non-declaration statement outside function body
//b:=30
var b = 30
func test() {
	a = 20
	fmt.Println("in test, a = ", a)
}
func main() {
	test()
	fmt.Println("a = ", a)
}