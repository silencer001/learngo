package dp

import (
	"fmt"
	"testing"
)

func TestPack01(t *testing.T) {
	iterms := []int{10, 4, 5, 6, 7, 3, 1}
	fmt.Println(Pack01(iterms, 9))
}

func TestPack01Optimized(t *testing.T) {
	iterms := []int{10, 4, 5, 6, 7, 3, 1}
	fmt.Println(Pack01Optimized(iterms, 9))
}

func TestPack02(t *testing.T) {
	iterms := []int{10, 4, 5, 6, 7, 3, 1}
	fmt.Println(pack02(iterms, 9))
}

func TestPack03(t *testing.T) {
	iterms := []int{10, 4, 5, 6, 7, 3, 1}
	fmt.Println(pack03(iterms, 9))
}

func TestPack04(t *testing.T) {
	iterms := []int{10, 4, 5, 6, 7, 3, 1}
	fmt.Println(pack04(iterms, 9))
}
