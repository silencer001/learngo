package mysort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	a := []int{2, 3, 1, 7, 5, 6, 10, 5, 5}
	fmt.Println(SelectSort(a))
}
func TestMergeSort(t *testing.T) {

	a := []int{2, 3, 1, 7, 5, 6, 10, 5, 5}
	fmt.Println(MergeSort(a))
}

func TestQuikSort(t *testing.T) {

	a := []int{2, 3, 1, 7, 5, 6, 10, 5, 5}
	fmt.Println(QuikSort(a))
}
