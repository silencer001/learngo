package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) < 2 {
		fmt.Println("useage: xxx file")
		return
	}

	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat: ", err)
		return
	}
	fmt.Printf("file info: %+v\n", info)
	fmt.Printf("file name:%s \n", info.Name()) //不包含路径
	fmt.Printf("file size %d\n", info.Size())
}
