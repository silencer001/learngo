package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"

	//1)解析规则，解析正则表达式，成功返回regexp.Regexp解释器
	reg1 := regexp.MustCompile(`a.c`)
	/*
		if reg1 == nil {//解析失败，panic
			fmt.Println(recover)
			return
		}*/
	//2)根据规则提取关键信息
	res1 := reg1.FindAllStringSubmatch(buf, 2)
	fmt.Println("res1 = ", res1)
}
