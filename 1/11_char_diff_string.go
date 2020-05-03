package main

import "fmt"
import "strings"

func main() {
	//var ch byte
	var str string

	//字符
	//1、单引号
	//2、往往都只有一个字符，转义字符除外'\n'
	//ch = 'a'
	//字符串
	//2、双引号
	//2、字符串有1个或多个字符组成
	//3、字符串都是隐藏了一个结束符'\0'
	str = "a啊"
	fmt.Printf("len(str) = %d, malloc(str) = %d", len(str), strings.Count(str, ""))

	//string类型可以用切片来引用？
	str = "hello go"
	fmt.Printf("str[0] = %c, str[1]= %c\n", str[0], str[1])
	//编译报错：cannot assin to str[0]，字符串字面量无法更改
	//str[0] = 'n'
}
