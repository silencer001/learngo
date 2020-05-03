package main

import "fmt"

func main() {
	src := []int{1,2,3,4}
	//dest := []int{6,6,6}
	//var dest []int
	dest := make([]int,1,10)
	//copy取两者len最小的那个
	copy(dest,src)
	fmt.Println("dest = ",dest)
}