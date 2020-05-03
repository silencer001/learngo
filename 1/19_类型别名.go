package main

import "fmt"

func main() {
	//给int64起一个别名叫做bigint
	type bigint int64

	var a bigint //等价于int64 a
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)
	var b long = 11
	var c char = 'a'
	fmt.Printf("b = %v, c = %v\n", b, c)
}
