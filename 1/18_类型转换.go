package main

import "fmt"

func main() {
	//不能转换的类型，就称为不兼容类型
	var flag bool
	flag = true
	//bool类型不能转换为整型
	//fmt.Printf("flag = %d\n", int(flag))
	fmt.Printf("flag = %t\n", flag)

	//error：cannot use 0 as type bool in assignment
	//整形也不能转换为bool
	//flag = 0
	//fmt.Printf("flag = %t\n", flag)

	//兼容类型
	var ch byte
	ch = 'a'
	//cannot use ch (type byte) as type int in assignment
	//var t int = ch
	var t int = int(ch) //类型转换，把ch的值取出来转换为int，再给t赋值
	fmt.Printf("t = %d\n", t)

}
