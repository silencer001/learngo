package calc

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	fmt.Printf("2+4*5+3=")
	res := InputExpress("2+4*5+3+1")
	fmt.Printf("%d\n", res)
}
