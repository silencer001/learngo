package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始计时")
	//2s后产生一个事件，
	//<-time.After(2 * time.Second)
	c := time.After(10 * time.Second)
	fmt.Println("start")
	<-c
	fmt.Println("时间到")
}

func main01() {
	//延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}
