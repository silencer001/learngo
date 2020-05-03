package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Append转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 0xff, 16)
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println("slice = ", string(slice))

	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)

	str = strconv.Itoa(666)
	fmt.Println("str = ", str)
	//字符串转其他类型
	flag, _ := strconv.ParseBool("false")
	fmt.Println("flag = ", flag)

	//字符串转换为整形
	i, _ := strconv.Atoi("12345")
	fmt.Println("i= ", i)
}
