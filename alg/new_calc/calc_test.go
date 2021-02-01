package calc

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	fmt.Printf("1+2+3=%d\n", Calc("1+2+3"))
	fmt.Printf("1+(1+2+3)/2=%d\n", Calc("1+(1+2+3)/2"))
	fmt.Printf("-1+2=%d\n", Calc("-1+2"))
}
