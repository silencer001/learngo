package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()
	fmt.Println("are u ok?")
	os.Stdout.WriteString("are u ok?\n")
	var a int
	os.Stdin.Close()
	fmt.Println("请输入a:")
	fmt.Scan(&a)
	fmt.Println("a=", a)
}
