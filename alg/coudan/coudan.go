package coudan

import "fmt"

/*mz为2，5，7，money*/
/*切片不能声明为const类型*/
var mz []int = []int{2, 5, 7}

//这个函数用的是回溯算法，实际上是深度递归，得出来的不是最优解
//也就是说动态规划不适合用递归啊？
func coudan(money int) int {
	/*假设money为27，27/2=13，如果14步都无法凑齐，说明无解*/
	maxstep := money/mz[0] + 1
	visited := make([][]bool, maxstep)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, 27)
	}
	var isFound bool = false
	needstep := 0
	type sugar func(int, int)
	var recall sugar
	recall = func(n int, curmoney int) {
		if n == maxstep || isFound { //not found
			return
		}
		if curmoney == money {
			isFound = true
			needstep = n
			return
		}
		if visited[n][curmoney] == true {
			return
		}
		visited[n][curmoney] = true
		if curmoney+mz[0] <= money {
			recall(n+1, curmoney+mz[0])
		}
		if curmoney+mz[1] <= money {
			recall(n+1, curmoney+mz[1])
		}
		if curmoney+mz[2] <= money {
			recall(n+1, curmoney+mz[2])
		}
	}
	recall(0, 0)
	return needstep
}

func coudan1(money int) int {
	states := make([][]bool, money/mz[0]+1)
	for i, _ := range states {
		states[i] = make([]bool, money+1)
	}
	states[0][2] = true
	states[0][5] = true
	states[0][7] = true
	isFound := false
	var step int = 1
	for step = 1; step < len(states); step++ {
		for i := 0; i <= money; i++ {
			if states[step-1][i] {
				for j := 0; j < len(mz); j++ {
					if i+mz[j] < money {
						states[step][i+mz[j]] = true
					}
					if i+mz[j] == money {
						isFound = true
						states[step][i+mz[j]] = true
						break
					}
				}
			}
			if isFound {
				break
			}
		}
		if isFound {
			break
		}
	}
	if !isFound {
		return -1
	}
	m := money
	fmt.Printf("money:\n")
	for i := step; i >= 1; i-- {
		for j := 0; j < len(mz); j++ {
			if states[i-1][m-mz[j]] {
				fmt.Printf("%d ", mz[j])
				m = m - mz[j]
				break
			}
		}
	}
	fmt.Println(m)
	fmt.Printf("\n")
	return step + 1
}

/*使用一个state数组，len为money+1，每个元素下标代表money，值代表最小的凑整个数
初始化均为-1，求state[money]的值
state[money] = min(state[money-mz[0]], state[money-mz[1]],state[money-mz[2]])
从小到大开始求每个index的最小凑单值,state[1]=-1,state[2]=1,state[3]=-1,state[4]=2,state[5]=1,state[6]=3,state[7]=1
*/
func coudan2(money int) int {
	state := make([]int, money+1)
	for i := 0; i < money+1; i++ {
		state[i] = -1
	}
	state[0] = 0                   //0用0个面值即可凑整完成
	for i := 1; i < money+1; i++ { //求每个money的凑整最小个数
		for j := 0; j < len(mz); j++ { //遍历每个面值
			if i-mz[j] >= 0 && state[i-mz[j]] != -1 { //i-mz[j]作为money，他的凑单最小值必须不为-1，即是有效值
				if state[i] == -1 || state[i-mz[j]]+1 < state[i] {
					state[i] = state[i-mz[j]] + 1
				}
			}
		}
	}
	if state[money] == -1 {
		return -1 //没找到
	}
	/*找出具体的凑整方案*/
	fmt.Printf("money:\n")
	for m := money; m > 0; {
		var min int = -1
		var minIndex int
		for _, v := range mz {
			if state[m-v] != -1 {
				if min == -1 || state[m-v] < min {
					minIndex = m - v
					min = state[m-v]
				}
			}
		}
		fmt.Printf("%d ", m-minIndex)
		//fmt.Printf("\ndebug %d %d\n", minIndex, state[minIndex])
		m = minIndex
	}
	fmt.Printf("\n")
	return state[money]
}
