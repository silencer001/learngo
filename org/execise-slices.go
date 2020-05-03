package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for index := range ret {
		ret[index] = make([]uint8, dy)
	}
	for x := range ret {
		for y := range ret[x] {
			//ret[x][y] = uint8(x) * uint8(y)
			switch {
			case y%15 == 0:
				ret[x][y] = uint8(x ^ y)
			case y%3 == 0:
				ret[x][y] = uint8(x ^ y)
			case y%5 == 0:
				ret[x][y] = uint8(x * y)
			}
		}
	}
	return ret
}
func main() {
	pic.Show(Pic)
}
