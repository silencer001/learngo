package main

import "fmt"

func main() {
	//自动推导类型
	a1 := [4]int{1, 2, 3, 4}
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 ：", s1)
	fmt.Println("a1 ：", a1)
	fmt.Printf("s1 type:%T, a1 type:%T\n", s1, a1)

	//借助make函数
	s2 := make([]int, 5, 10)
	fmt.Printf("len(s2)=%d, cap(s2)=%d\n", len(s2), cap(s2))
	s3 := make([]int, 5)
	fmt.Printf("len(s3)=%d, cap(s3)=%d\n", len(s3), cap(s3))
}
func main01() {
	//切片和数组的区别
	//数组[]里面的长度是固定的常量
	a := [5]int{}
	fmt.Printf("len(a)=%d, cap(a)=%d\n", len(a), cap(a))

	//切片，[]里面为空，或者为...切片的长度容量不固定
	s := []int{}
	fmt.Printf("len(s)=%d, cap(s)=%d\n", len(s), cap(s))
	s = append(s, 11)
	fmt.Printf("len(s)=%d, cap(s)=%d\n", len(s), cap(s))

}
