package main

import (
	"fmt"
)

func main() {
	var array = [3]int{1, 2, 3} //array can't use append api
	fmt.Printf("%s len=%d cap=%d %v\n", "array", len(array), cap(array), array)
	//array = append(array, 7)
	var slice = []int{4, 5, 6}
	fmt.Printf("%s len=%d cap=%d %v\n", "slice", len(slice), cap(slice), slice)
	slice = append(slice, 7)
	fmt.Printf("%s len=%d cap=%d %v\n", "slice after append", len(slice), cap(slice), slice)
}
