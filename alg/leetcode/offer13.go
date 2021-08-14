package leetcode

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格
（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思考：
通过画图，发现是从0,0开始的一层一层的范围，看起来像是bfs搜索，bfs搜索考虑用队列*/
func movingCount(m int, n int, k int) int {
	state := make([][]bool, m)
	for i, _ := range state {
		state[i] = make([]bool, n)
	}
	type point struct {
		x int
		y int
	}
	var queue []point
	var res int
	queue = append(queue, point{0, 0})
	count := func(x int) int {
		res := 0
		for x > 0 {
			res += (x % 10)
			x = x / 10
		}
		return res
	}
	move := func(p point) bool {
		{
			counter := count(p.x) + count(p.y)
			if counter > k {
				return false
			} else {
				return true
			}
		}
	}
	addqueue := func(p point) {
		if p.x+1 < m && !state[p.x+1][p.y] {
			queue = append(queue, point{p.x + 1, p.y})
		}
		if p.y+1 < n && !state[p.x][p.y+1] {
			queue = append(queue, point{p.x, p.y + 1})
		}
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		state[p.x][p.y] = true
		if move(p) {
			res++
			addqueue(p)
		} else {

		}
	}
	return res
}
