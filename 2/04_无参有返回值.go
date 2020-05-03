package main

import "fmt"
//只有一个返回值
//需要return关键字返回
func myfunc01() int {
	return 666
}
//给返回值起一个变量名，go推荐语法
func myfunc02() (result int) {
	result = 6666
	return result//必须加return，但可以不加result
}

func myfunc03() (result int) {
	result = 66666
	return //必须加return，但可以不加result
}
func main() {
	i:=myfunc01()
	fmt.Println(i)
	j:=myfunc02()
	fmt.Println(j)
	fmt.Println(myfunc03())
}