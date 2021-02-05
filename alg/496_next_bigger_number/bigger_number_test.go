package bigger_number

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := NextGreaterElement(nums1, nums2)
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(res)
}
