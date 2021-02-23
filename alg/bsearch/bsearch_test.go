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
