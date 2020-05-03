package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("this is a NewTask")
		time.Sleep(time.Second)
	}
}

func main() {
	go NewTask()
	for {
		fmt.Println("this is main go routine")
		time.Sleep(time.Second)
	}
}
