package findv

import (
	"fmt"
	"testing"
)

func TestFindV(t *testing.T) {
	a := []int{2, 6, 9, 10, 3, 4, 7, 1, 15, 5}
	fmt.Println(FindV(a, 10))
	fmt.Println(FindV(a, 7))
	fmt.Println(FindV(a, 15))
}
