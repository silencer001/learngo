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

func TestCountSort(t *testing.T) {

	a := []int{2, 3, 1, 1, 5, 4, 4, 5, 5, 0, 4}
	fmt.Println(CountSort(a))
}

func TestFindk(t *testing.T) {

	a := []int{2, 3, 1, 7, 5, 6, 10, 5, 5}
	fmt.Println(FindK(a, 7))
}
