package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsdg  34.1234 1.23 7. 8.99 lsdjalsdf 6.66"
	reg := regexp.MustCompile(`[0-9]+\.\d`)
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(res)
}
