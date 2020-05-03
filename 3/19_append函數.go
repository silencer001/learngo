package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)

	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	s2 = append(s2, 3)
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))

}
