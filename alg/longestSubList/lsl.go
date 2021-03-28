package longestSubList

/*一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？
比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长
递增子序列的长度是 4*/

/*思考
从第一个元素开始排序不一定是最长子串，无法使用贪心算法
考虑使用动态规划，先求状态转移方程
lsl(i)=MAX(if(a[i]>a[i-1]lsl[i-1]+1;同理lsl[i-2],lsl[i-3]))
*/

func lsl(a []int) int {
	sl := make([]int, len(a))
	//设定初始值，每个状态的最长递增序列最小为1
	for i, _ := range sl {
		sl[i] = 1
	}
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[j] < a[i] && sl[i] < sl[j]+1 {
				sl[i] = sl[j] + 1
			}
		}
	}
	max := 1
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return max
}
