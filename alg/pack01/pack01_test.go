package pack01

import (
	"fmt"
	"testing"
)

func TestPack01(t *testing.T) {
	iterms := []int{10, 14, 20, 25, 30, 1, 2}
	fmt.Println(Pack01(iterms, 35))
}
