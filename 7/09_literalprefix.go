package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`.abc`)
	fmt.Println(reg.LiteralPrefix())
}
