package main

import "fmt"

func main() {
	s := "王思聪"
	if s == "王思聪" {
		fmt.Printf("左手一个妹子，右手一个大妈\n")
	} else {
		fmt.Println("好好写代码")
	}

	//if支持1个初始化语句
	if a := 10; a == 10 {
		fmt.Printf("a=%d\n", a+1)
	}
	//a的作用域在if块种
	//fmt.Printf("a = %d\n", a-1)

}
