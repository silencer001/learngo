package main

import "fmt"
func main() {
	var ch byte
	ch =97
	fmt.Println("ch= ", ch)
	fmt.Printf("ch= %c， or ch= %d\n", ch, ch)

	ch = 'a'//字符必须是单引号
	fmt.Printf("%c, %d\n", ch, ch)
	//大写转小写，小写转大写
	fmt.Printf("大写：%d, 小写：%d\n", 'A', 'a')
	ch -=32
	fmt.Printf("%c, %d\n", ch, ch)
}