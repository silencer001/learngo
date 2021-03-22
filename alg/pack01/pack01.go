package pack01

import "fmt"

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

func Pack01Dynamic(items []int, limit int) int {
	len := len(items)
	max := 0
	//packed := make([]bool, len)
	var packed []int
	isRepeat := make([][]bool, len)
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
