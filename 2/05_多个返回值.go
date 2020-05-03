package main

import "fmt"

func myfunc01() (i int, j int, m int) {
	return 1, 2, 3
}

//只有一个返回值时可以省略名称和小括号
func myfunc02(istring int) int {
	if istring == 1 {
		return 0
	} else {
		return 1
	}
}
func main() {
	//multiple-value myfunc01() in single-value context
	// for i,v := range myfunc01() {
	// 	fmt.Println(i,v)
	// }
	fmt.Println(myfunc01())
	fmt.Println(myfunc02(0))
}
