package main

import "fmt"

func MyFunc01(a int, b int) { //固定参数

}

//MyFunc02:...int类型这样的类型，...type不定参数类型
func MyFunc02(args ...int) {
	fmt.Println("len(args) = ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}
	for i, v := range args {
		fmt.Printf("args[%d] = %d\n", i, v)

	}
}
func testFunc02(args2 ...int) {
	//全部元素传递给myFunc02
	MyFunc02(args2...)
	//只想把后2个参数传递给另外一个函数
	MyFunc02(args2[1:3]...)
	MyFunc02(args2[1:]...)
	MyFunc02(args2[:3]...)
}
func myfunc03(a string, args ...int) {

}

//syntax error: cannot use ... with non-final parameter args
// func myfunc04(args ...int, b int) {

// }
func main() {
	MyFunc01(1, 2)
	MyFunc02(1)
	MyFunc02(1, 2, 3)
	testFunc02(4, 3, 2, 1)
}
