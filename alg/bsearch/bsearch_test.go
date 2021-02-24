package bsearch

import (
	"fmt"
	"testing"
)

func TestBsearch(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 20, 30, 40}
	fmt.Println("a:=", a)
	fmt.Println("value 8 index:", Bsearch(a, 8))
	fmt.Println("value 8 index:", BsearchRecurs(a, 8))
}

func TestBsearchFirst(t *testing.T) {

	a := []int{0, 1, 2, 3, 4, 5, 5, 5, 6, 7, 7, 9, 20, 30, 40}
	fmt.Println("a:=", a)
	fmt.Println("First 7 index:", BsearchFirst(a, 7))
	fmt.Println("First 5 index:", BsearchFirst(a, 5))
}

func TestBsearchLow(t *testing.T) {

	a := []int{2, 3, 4, 5, 5, 5, 6, 7, 7, 9, 20, 30, 40}
	fmt.Println("a:=", a)
	fmt.Println("First 7 index:", BsearchLow(a, 7))
	fmt.Println("First 8 index:", BsearchLow(a, 8))
	fmt.Println("First 1 index:", BsearchLow(a, 1))
}
