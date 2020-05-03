package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max]取下标从low开始的元素，len=high-low, cap=max-low
	s1 := array[:5] //不指定容量时，容量和底层数组的容量相同
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	a1 := make([]int, 10, 20)
	s2 := a1[:10] //cap(s2) = cap(a1)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	a1[1] = 1
	fmt.Println("s2:", s2)
	s3 := a1[3:6:7]
	s3[2] = 3
	fmt.Println("s3", s3)
	fmt.Println("s2:", s2)
}
