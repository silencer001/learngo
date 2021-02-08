package selectSort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{2, 3, 1, 7, 6}
	fmt.Println(Sort(a))
}
