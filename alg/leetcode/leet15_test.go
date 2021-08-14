package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	ans := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for _, v := range ans {
		fmt.Println(v)
	}
}
