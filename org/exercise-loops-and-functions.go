package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var res float64 = 1
	var z float64 = 0
	if x == 1 {
		return 1
	}

	for math.Abs(res-z) > 0.00000001 {
		z = res
		res = z - (math.Pow(z, 2)-x)/(2*z)
		fmt.Println(res)
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
}
