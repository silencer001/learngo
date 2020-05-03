package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	fmt.Println("cmd line len :", len(list))
	for _, value := range os.Args {
		fmt.Println("param:", value)
	}
}
