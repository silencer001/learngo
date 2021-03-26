package levenshtein

/*求把一个字符串变成另一个字符串，需要的最少编辑次数*/

/*
莱温斯坦算法允许添加、删除、修改三种操作，使得A字符串变为B字符串，求最少的编辑次数
思路：先使用回溯法？递归遍历所有的可能，记录最小编辑次数
*/
func levenshtein1(a string, b string) int {
	//用二维bool数组会有bug，visted[i][j]在i，j时会有多种结果，保存最小的，其他大的遍历没有意义了
	visited := make([][]int, len(a))
	for i, _ := range visited {
		visited[i] = make([]int, len(b))
		for j, _ := range visited[i] {
			visited[i][j] = 0xffffffff //设定一个最大值
		}
	}

	min := -1
	type recall func(ai, bj, curdst int)
	var levendist recall
	levendist = func(ai, bj, curdst int) {
		if ai == len(a) || bj == len(b) {
			curdst += (len(b) - bj) + (len(a) - ai)
			if min == -1 || curdst < min {
				min = curdst
			}
			return
		}
		if curdst > visited[ai][bj] {
			return
		}
		visited[ai][bj] = curdst
		if a[ai] == b[bj] {
			levendist(ai+1, bj+1, curdst)
		} else {
			levendist(ai+1, bj, curdst+1)
			levendist(ai, bj+1, curdst+1)
			levendist(ai+1, bj+1, curdst+1)
		}
		return
	}
	levendist(0, 0, 0)
	return min
}
