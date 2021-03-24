package matrix

/*
有一个 n 乘以 n 的矩阵 w[n][n]。矩阵存储的都是正整数。棋子起始位置在左上
角，终止位置在右下角。我们将棋子从左上角移动到右下角。每次只能向右或者向下移动一
位。从左上角到右下角，会有很多不同的路径可以走。我们把每条路径经过的数字加起来看
作路径的长度。那从左上角移动到右下角的最短路径长度是多少呢？
*/

/*
从m[0][0]到m[i][j]的最短路径问题
第一种就是使用回溯法，找出所有的路径，然后记录最小的cost
优化点是有回溯中有已经访问过的路径，使用cache进行记录cost
*/
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func findCostMin1(m [][]int, i int, j int) int {
	min := 0
	type recall func(i int, j int, cost int)
	var findmin recall
	findmin = func(i int, j int, cost int) {
		if i == len(m)-1 && j == len(m[0])-1 {
			if min == 0 || cost < min {
				min = cost
			}
			return
		}
		if i < len(m)-1 {
			findmin(i+1, j, cost+m[i][j])
		}
		if j < len(m[0])-1 {
			findmin(i, j+1, cost+m[i][j])
		}
	}
	findmin(0, 0, 0)
	return min + m[i][j]
}

/*
使用动态规划法，首先推导状态转移方程
MIN(i,j) = min(MIN(i-1,j),MIN(i,j-1))+m[i][j]
同样是递归，但这个复杂度应该比回溯要好，因为没有遍历所有的路径，而是倒推最优解？
*/
func findCostMin2(m [][]int, i int, j int) int {
	if i == 0 && j == 0 {
		return m[0][0]
	}
	if i == 0 {
		return m[i][j] + findCostMin2(m, 0, j-1)
	}
	if j == 0 {
		return m[i][j] + findCostMin2(m, i-1, 0)
	}
	return min(findCostMin2(m, i, j-1), findCostMin2(m, i-1, j)) + m[i][j]
}

/*
对findCostMin2进行优化剪枝操作
*/
func findCostMin3(m [][]int, i int, j int) int {
	visted := make([][]int, i+1)
	for i, _ := range visted {
		visted[i] = make([]int, j+1)
	}
	for i := 0; i < len(visted); i++ {
		for j := 0; j < len(visted[i]); j++ {
			visted[i][j] = -1
		}
	}
	type recall func(int, int) int
	var findCost recall
	findCost = func(i int, j int) int {
		if i == 0 && j == 0 {
			return m[0][0]
		}
		if visted[i][j] != -1 {
			return visted[i][j]
		}

		if i == 0 {
			visted[i][j] = m[i][j] + findCost(0, j-1)
		} else if j == 0 {
			visted[i][j] = m[i][j] + findCost(i-1, 0)
		} else {
			visted[i][j] = min(findCost(i, j-1), findCost(i-1, j)) + m[i][j]
		}
		return visted[i][j]
	}
	return findCost(i, j)
}

/*使用二维的状态转移表，每个i,j代表对应的点，值代表最小cost*/
func findCostMin4(m [][]int, i int, j int) int {
	state := make([][]int, len(m))
	for n, _ := range state {
		state[n] = make([]int, len(m[n]))
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			state[i][j] = -1
		}
	}
	//对state赋起始值
	state[0][0] = m[0][0]
	//对state的第一列进行赋值
	for i := 1; i < len(m); i++ {
		state[i][0] = state[i-1][0] + m[i][0]
	}
	//对state的第一行进行赋值
	for i := 1; i < len(m[0]); i++ {
		state[0][i] = state[0][i-1] + m[0][i]
	}
	//状态转移表，求每个点的最小cost
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			state[i][j] = min(state[i-1][j], state[i][j-1]) + m[i][j]
		}
	}
	return state[i][j]
}
