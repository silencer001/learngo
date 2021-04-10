package dp

import (
	"fmt"
)

func Pack01(items []int, limit int) int {
	len := len(items)
	max := 0
	//packed := make([]bool, len)
	var packed []int
	type xxx func(int, int, []int)
	var pack xxx
	pack = func(cur int, nowweight int, flag []int) {
		if cur == len || max == limit {
			if nowweight > max { //max的赋值放在cur == len即遍历到最后一层时,nowweight必定小于limit
				max = nowweight
				packed = flag
			}
			return
		}
		pack(cur+1, nowweight, flag)       //不把当前item放入pack
		if nowweight+items[cur] <= limit { //把当前item放入pack，进行剪枝操作
			flag1 := append(flag, cur)
			pack(cur+1, nowweight+items[cur], flag1)
		}
	}
	flag := []int{}
	pack(0, 0, flag)
	fmt.Println(packed)
	return max
}

func Pack01Optimized(items []int, limit int) int {
	len := len(items)
	max := 0
	//packed := make([]bool, len)
	var packed []int
	isRepeat := make([][]bool, len)
	/*到每一层遍历时，nowweight最多只有limit+1个状态*/
	for i := 0; i < len; i++ {
		isRepeat[i] = make([]bool, limit+1)
	}
	type recallFunc func(int, int, []int)
	var pack recallFunc
	pack = func(cur int, nowweight int, flag []int) {
		if cur == len || max == limit {
			if nowweight > max { //max的赋值放在cur == len即遍历到最后一层时,nowweight必定小于limit
				max = nowweight
				packed = flag
			}
			return
		}
		if isRepeat[cur][nowweight] {
			return //该场景已经被遍历过了
		}
		isRepeat[cur][nowweight] = true
		pack(cur+1, nowweight, flag)       //不把当前item放入pack
		if nowweight+items[cur] <= limit { //把当前item放入pack，进行剪枝操作
			flag1 := append(flag, cur)
			pack(cur+1, nowweight+items[cur], flag1)
		}
	}
	flag := []int{}
	pack(0, 0, flag)
	fmt.Println(packed)
	return max
}

/*推导过程：
1、有n个元素，选择其中m个，达到最大值，不超过limit
2、推导公式：假设遍历第n个元素时的最大值是MAXn，第n-1步的最大值是MAXn-1
MAXn=max(MAXn-1,MAXn-1 + iterms[n])，其中MAXn-1 + iterms[n]必须小于limit
3、起始值MAX0=0
*/
func pack02(iterms []int, limit int) int {
	states := make([]int, len(iterms)+1)
	for i, _ := range states {
		states[i] = 0
	}
	states[0] = 0
	var i int
	for i = 1; i < len(states); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if states[j]+iterms[i] <= limit {
				if states[j]+iterms[i] > max {
					max = states[j] + iterms[i]
				}
			} else {
				if states[j] > max {
					max = states[j]
				}
			}
		}
		states[i] = max
		if max == limit {
			break
		}
	}
	return states[i]
}

/*
推导过程：
以树的形式推导每一步骤的可能重量
		f(0,0)
		/     \
	f(1,0)		f(1,2)
	/	\		/	  \
f(2,0)	f(2,4) f(2,4) f(2,6)
每一层的重量可能为0,4,6
申请一个n X limit的二维bool类型数组
每一行代表一层步骤，每一层中为true的下标代表这一步骤可以到达的weight
如果可以达到limit，就说明找到了，此时可以退出，没有找到就继续下一行

最后在最后一行从后往前遍历，找到的第一个为true的下标就代表可以到达的最大重量
*/
func pack03(iterms []int, limit int) int {
	states := make([][]bool, len(iterms)+1)
	for i, _ := range states {
		states[i] = make([]bool, limit+1)
	}
	states[0][0] = true
	for i := 1; i <= len(iterms); i++ {
		w := iterms[i-1]
		for j := 0; j <= limit; j++ {
			if states[i-1][j] {
				states[i][j] = true
				if j+w <= limit {
					states[i][j+w] = true
				}
			}
		}
	}
	for i := limit; i >= 0; i-- {
		if states[len(iterms)][i] {
			return i
		}
	}
	return -1
}

func pack04(iterms []int, limit int) int {
	states := make([]bool, limit+1)
	states[0] = true
	for i := 1; i <= len(iterms); i++ {
		//for j:=0;j<=limit-iterms[i-1];j++ { 从小到大遍历会j+iterms[i-1]变为true，然后又遍历了一次
		for j := limit - iterms[i-1]; j >= 0; j-- {
			if states[j] {
				states[j+iterms[i-1]] = true
			}
		}
	}
	for i := limit; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return -1
}
