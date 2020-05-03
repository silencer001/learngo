package main

import (
	"fmt"
	"time"
)

func main() {
	//for后面不写任何东西，循环条件永远为真，死循环
	i := 0
	for {
		i++
		time.Sleep(time.Second)

		if i == 5 {
			//break//跳出最近的内循环
			continue //跳出本次循环
		}
		fmt.Println("i= ", i)
	}
}
