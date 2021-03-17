package queen8

import "fmt"

func Queen8() []int {
	queen := make([][]int, 8)
	for i := 0; i < len(queen); i++ {
		queen[i] = make([]int, 8)
	}
	//从start列开始深度遍历
	conflict := func(column, row int) bool {
		for i := 0; i < column; i++ {
			if queen[i][row] == 1 {
				return true
			}
		}
		for i := 0; i < 8; i++ {
			if i != row && queen[column][i] == 1 {
				return true
			}
		}
		for i, j := column-1, row-1; i >= 0 && j >= 0; i-- {
			if queen[i][j] == 1 {
				return true
			}
			j--
		}
		for i, j := column-1, row+1; i >= 0 && j < 8; i-- {
			if queen[i][j] == 1 {
				return true
			}
			j++
		}
		return false
	}
	type sugarfunc func(int) bool
	var search sugarfunc
	search = func(start int) bool {
		for i := 0; i < 8; i++ {
			fmt.Println("row:", i)
			queen[start][i] = 1
			//验证是否有冲突
			if conflict(start, i) {
				queen[start][i] = 0
				continue
			} else {
				if search(start + 1) {

					return true
				} else {
					queen[start][i] = 0
					continue
				}
			}
		}
		fmt.Println("not found at column:", start)
		return false
	}
	res := make([]int, 8)
	if !search(0) {
		return nil
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if queen[i][j] == 1 {
				res[i] = j
				break
			}
		}
	}

	return res
}
