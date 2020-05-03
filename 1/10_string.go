package main

import "fmt"

func main() {
	var str1 string
	str1 = "abc"
	fmt.Println("str1= ", str1)

	//自动推导类型
	str2 := "mike\n"
	fmt.Printf("str 类型是 %T\n", str2)

	//内建函数，len()可以检测字符串的长度有多少个字符
	fmt.Println("len(str2) = ", len(str2))
}
