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

	min := 0xfffffffff
	type recall func(ai, bj, curdst int)
	var levendist recall
	levendist = func(ai, bj, curdst int) {
		if ai == len(a) || bj == len(b) {
			curdst += (len(b) - bj) + (len(a) - ai)
			if curdst < min {
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

/*
有重复子问题，首先考虑动态规划
首先求状态转移方程
mindst(i,j) = min(mindst(i,j-1)+1, mindst(i-1,j)+1,mindst(i-1,j-1)或mindst(i-1,j-1)+1)
使用state矩阵来表示
	A	C	B	D
A	0	1	2	3
D	1   1	2	2
C	2	1	2	3
B	3	2	1	2
*/

func levenshtein2(a string, b string) int {
	state := make([][]int, len(a))
	for i, _ := range state {
		state[i] = make([]int, len(b))
	}
	/*有bug
	for i := 1; i < len(a); i++ {
		state[i][0] = state[i-1][0] + 1
	}
	for j := 1; j < len(b); j++ {
		state[0][j] = state[0][j-1] + 1
	}*/
	if a[0] == b[0] {
		state[0][0] = 0
	} else {
		state[0][0] = 1
	}
	for i := 1; i < len(a); i++ {
		if a[i] == b[0] {
			state[i][0] = i
		} else {
			state[i][0] = state[i-1][0] + 1
		}
	}
	for j := 1; j < len(b); j++ {
		if b[j] == a[0] {
			state[0][j] = j /*差距j个字符，因此距离是j*/
		} else {
			state[0][j] = state[0][j-1] + 1
		}
	}
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				state[i][j] = min(state[i-1][j]+1, state[i][j-1]+1, state[i-1][j-1])
			} else {
				state[i][j] = min(state[i-1][j]+1, state[i][j-1]+1, state[i-1][j-1]+1)
			}
		}
	}
	return state[len(a)-1][len(b)-1]
}

func min(i, j, k int) int {
	min := i
	if j < min {
		min = j
	}
	if k < min {
		min = k
	}
	return min
}

/*求公共最长子串的长度，允许使用删除、添加两种操作
实际上仍然是求最小的修改次数，然后用a b中其中较长的字符串长度减去最小修改次数，就是最长子串长度了
这个想法不正确，如果是添加方法，应该是用添加后的字符串长度减去最小修改次数

每一步的动作，对下一步有影响，所以无法使用贪心算法
使用回溯算法可行，效率太低
考虑使用动态规划算法，首先求状态转移方程
a[i-2]  a[i-1] a[i]
b[j-2]  b[j-1] b[j]
如果a[i]和b[j]相等，那么
lcs(i,j)=MAX(lcs(i-1,j),lcs(i,j-1),lcs(i-1,j-1)+1)
如果a[i]和b[j]不等，那么
lcs(i,j)=MAX(lcs(i-1,j),lcs(i,j-1),lcs(i-1,j-1))
*/
func lcs(a, b string) int {
	lcs := make([][]int, len(a))
	for i, _ := range lcs {
		lcs[i] = make([]int, len(b))
	}
	if a[0] == b[0] {
		lcs[0][0] = 1
	} else {
		lcs[0][0] = 0
	}
	for i := 1; i < len(a); i++ {
		if a[i] == b[0] {
			lcs[i][0] = 1
		} else {
			lcs[i][0] = lcs[i-1][0]
		}
	}
	for j := 1; j < len(b); j++ {
		if b[j] == a[0] {
			lcs[0][j] = 1
		} else {
			lcs[0][j] = lcs[0][j-1]
		}
	}
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1], lcs[i-1][j-1]+1)
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1], lcs[i-1][j-1])
			}
		}
	}
	return lcs[len(a)-1][len(b)-1]
}
func max(i, j, k int) int {
	max := i
	if j > max {
		max = j
	}
	if k > max {
		max = k
	}
	return max
}
