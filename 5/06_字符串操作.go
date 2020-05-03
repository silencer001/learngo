package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo"中是否包含"hello"
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	str := strings.Join(s, ", ")
	fmt.Println(str)

	//Index索引
	fmt.Println((strings.Index("abcdhello", "fhello")))

	buf := strings.Repeat("go", 3)
	fmt.Println("buf=", buf)

	//Split
	fmt.Println(strings.Split(str, ","))
	fmt.Println(strings.Split(str, ", "))

	//Trim去掉两头的字符
	buf = strings.Trim("        are u OK?     ", " ")
	fmt.Println("buf=", buf)

	s3 := strings.Fields("       are u OK?  ")
	for _, s := range s3 {
		fmt.Println(s)
	}
}
