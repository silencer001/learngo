package main

import (
	"fmt"
	"time"
)
//验证timer的周期性
func main() {
	timer := time.NewTimer(1*time.Second)
	for {
		//fatal error: all goroutines are asleep - deadlock!
		<-timer.C
		fmt.Println("时间到")
	}
}
func main01() {
	//创建一个定时器，设置时间为2s，2s后，向time通道写内容（当前时间）
	timer := time.NewTimer(2 *time.Second)
	fmt.Println("current time:", time.Now())
	t := <-timer.C
	fmt.Println("t=",t)
}