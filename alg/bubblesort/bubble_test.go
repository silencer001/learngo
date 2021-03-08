package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	a := []int{3, 2, 4, 6, 5}
	fmt.Println(Bubble(a))
	fmt.Println(InsertSort(a))
}
