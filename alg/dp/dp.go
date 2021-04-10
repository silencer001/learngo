package dp

import "fmt"

func myAtoi(s string) int {
	negflag := false
	start := false
	res := 0
	for _, c := range s {
		if !start && c == ' ' {
			continue
		}
		if !start && c == '-' {
			negflag = true
			start = true
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		start = true
		fmt.Println(int(c))
		res = res*10 + int(c) - 48
		if negflag && res > 231 {
			res = 231
			break
		}
		if res >= 231 {
			res = 230
			break
		}

	}
	if negflag {
		res *= -1
	}
	return res
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	isdup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		isdup[nums[i]]++
	}
	res = append(res, []int{})
	for v, num := range isdup {
		//fmt.Println(v, num)
		clen := len(res)
		fmt.Println("res:  ", res)
		for k := 0; k < num; k++ {
			for j := 0; j < clen; j++ {
				fmt.Println("j=", j, "   ", res[clen*k+j])
				n := make([]int, len(res[clen*k+j]))
				copy(n, res[clen*k+j])
				n = append(n, v)
				fmt.Println(n)
				res = append(res, n)
			}
		}

	}
	return res
}

/*
res:   [[]]
j= 0     []
[4]
j= 0     [4]
[4 4]
j= 0     [4 4]
[4 4 4]
j= 0     [4 4 4]
[4 4 4 4]
res:   [[] [4] [4 4] [4 4 4] [4 4 4 4]]
j= 0     []
[1]
j= 1     [4]
[4 1]
j= 2     [4 4]
[4 4 1]
j= 3     [4 4 4]
[4 4 4 1]
j= 4     [4 4 4 1]
[4 4 4 1 1]
[[] [4] [4 4] [4 4 4] [4 4 4 1] [1] [4 1] [4 4 1] [4 4 4 1] [4 4 4 1 1]]
*/
func subsetsWithDup1(nums []int) [][]int {
	res := make([][]int, 0)
	isdup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		isdup[nums[i]]++
	}
	res = append(res, []int{})
	for v, num := range isdup {
		//fmt.Println(v, num)
		len := len(res)
		fmt.Println("res:  ", res)
		for k := 0; k < num; k++ {
			for j := 0; j < len; j++ {
				fmt.Println("j=", j, "   ", res[len*k+j])
				n := append(res[len*k+j], v)
				fmt.Println(n)
				res = append(res, n)
			}
		}

	}
	return res
}

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*动态规划算法求解，先列出状态转移方程
MAX(n)=MAX(max(n-1)+1,max(n-2)+1,max(n-3)+1)//只有当n-3时的01个数加上第n个01个数小于limit时
*/
func findMaxForm(strs []string, m int, n int) int {
	state := make([][]int, m+1)
	for i, _ := range state {
		state[i] = make([]int, n+1)
		for j, _ := range state[i] {
			state[i][j] = -1
		}
	}
	state[0][0] = 0
	for _, str := range strs {
		str1 := 0
		for _, c := range str {
			if c == '1' {
				str1++
			}
		}
		str0 := len(str) - str1
		for i := 0; i <= m-str0; i++ {
			for j := 0; j <= n-str1; j++ {
				if state[i][j] != -1 && state[i][j]+1 > state[i+str0][j+str1] {
					state[i+str0][j+str1] = state[i][j] + 1
				}
			}
		}

		/*
			for i := m - str0; i >= 0; i-- {
				for j := n - str1; j >= 0; j-- {
					if state[i][j] != -1 && state[i][j]+1 > state[i+str0][j+str1] {
						state[i+str0][j+str1] = state[i][j] + 1
					}
				}
			}*/
	}
	max := 0
	for _, a := range state {
		for _, v := range a {
			if max < v {
				max = v
			}
		}
	}
	return max
}

/*
给定一个字符串 S，找出 S 中不同的非空回文子序列个数，并返回该数字与 10^9 + 7 的模。

通过从 S 中删除 0 个或多个字符来获得子序列。

如果一个字符序列与它反转后的字符序列一致，那么它是回文字符序列。

如果对于某个  i，A_i != B_i，那么 A_1, A_2, ... 和 B_1, B_2, ... 这两个字符序列是不同的。



示例 1：

输入：
S = 'bccb'
输出：6
解释：
6 个不同的非空回文子字符序列分别为：'b', 'c', 'bb', 'cc', 'bcb', 'bccb'。
注意：'bcb' 虽然出现两次但仅计数一次。
示例 2：

输入：
S = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
输出：104860361
解释：
共有 3104860382 个不同的非空回文子序列，对 10^9 + 7 取模为 104860361。


提示：

字符串 S 的长度将在[1, 1000]范围内。
每个字符 S[i] 将会是集合 {'a', 'b', 'c', 'd'} 中的某一个。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-different-palindromic-subsequences

*/
func countPalindromicSubsequences(S string) int {
	//元数组
	ch := []byte{'a', 'b', 'c', 'd'}
	slen := len(S)
	dp := make([][][]int, len(ch))
	for i, _ := range dp {
		dp[i] = make([][]int, slen)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, slen)
		}
	}

	//初始化每个字符的状态转移矩阵，单个字符的回文个数只有1
	for i := slen - 1; i >= 0; i-- {
		for j := i; j < slen; j++ {
			for k := 0; k < len(ch); k++ {
				c := ch[k]
				if i == j {
					if c != S[i] {
						dp[k][i][j] = 0
					} else {
						dp[k][i][j] = 1
					}
					continue
				}
				if c != S[i] {
					dp[k][i][j] = dp[k][i+1][j]
				} else if c != S[j] {
					dp[k][i][j] = dp[k][i][j-1]
				} else { //S[i]==S[j]==c
					if j == i {
						dp[k][i][j] = 1
					} else if j == i+1 { //相邻的两个字符相同
						dp[k][i][j] = 2
					} else {
						dp[k][i][j] = 2 + dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1]
					}
				}
			}
		}
	}
	//debug
	for i := 0; i < 4; i++ {
		//fmt.Println(ch[i], " : ", dp[i][0][slen-1])
	}
	return dp[0][0][slen-1] + dp[1][0][slen-1] + dp[2][0][slen-1] + dp[3][0][slen-1]
}

/*
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
思考1：为什么要用动态规划?
最长xx问题，但不是第一个回文字符是最长的，因此贪心算法不适用，回文会有很多重复解，有最优子结构
思考2：确定适用dp后，如何求状态转移方程?
f(n)=f(n-1) + {如果n有对应的字符，则为1，如果没有，则为0}
不合适，n对应的字符需要是非f(n-1)回文中的字符，考虑区间dp？
换种思路最长回文，应该是去掉最少的字符，像levenstan替换那样？即删掉最少个字符成为回文字符串?
没思路，就用区间dp
f(i,j)=f(i+1,j-1) + cost{if(s[i]==s[j])cost=2;else cost=0}
思考3：状态转移矩阵，考虑用i作为阶，j作为宽度len
  b  a  b  a  d
b 1
a	 1
b		1
a			1  1
b			   1
*/
func max(a, b, c int) int {
	k := a
	if k < b {
		k = b
	}
	if k < c {
		k = c
	}
	return k
}
func longestPalindrome(s string) string {
	slen := len(s)
	dp := make([][]int, slen)
	for i, _ := range dp {
		dp[i] = make([]int, slen)
	}
	for i := slen - 1; i >= 0; i-- {
		for j := i; j < slen; j++ {
			if i == j {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else {
				if dp[i+1][j-1] == 1 && s[i] == s[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			}
		}
	}

	for len := slen - 1; len > 0; len-- {
		for i := 0; i+len < slen; i++ {
			if dp[i][i+len] == 1 {
				return s[i : i+len+1]
			}
		}
	}
	return string(s[0])
}

/*求非连续的最长回文子串*/
func longestPalindrome2(s string) string {
	slen := len(s)
	dp := make([][]int, slen)
	for i, _ := range dp {
		dp[i] = make([]int, slen)
	}
	for i := slen - 1; i >= 0; i-- {
		for j := i; j < slen; j++ {
			if i == j {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = max(dp[i+1][j-1]+2, dp[i+1][j], dp[i][j-1])
				} else {
					dp[i][j] = max(dp[i+1][j-1], dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	ans := make([]byte, dp[0][slen-1])
	j := 0
	count := len(ans)
	for i := slen - 1; i > 0 && count > 0; i-- {
		if dp[0][i] > dp[0][i-1] {
			ans[j] = s[i]
			j++
			count -= 2
			continue
		}
		if count == 1 {
			ans[j] = s[i]
			count -= 1
		}
	}
	for i := len(ans) - 1; i >= len(ans)/2; i-- {
		ans[i] = ans[len(ans)-i-1]
	}
	return string(ans)
}

/*错误解法*/
func longestPalindrome1(s string) string {
	slen := len(s)
	if slen == 1 {
		return s
	}
	if slen == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[1:]
		}
	}
	dp := make([]int, slen)
	mid := slen / 2
	dp[mid] = 1
	if s[mid] == s[mid-1] {
		dp[mid-1] = 2
		dp[mid] = 2
	} else {
		dp[mid-1] = 1
	}
	for i := 1; i < slen-mid; i++ {
		if s[mid+i] == s[mid-i] {
			dp[mid+i] = dp[mid+i-1] + 2
		} else {
			dp[mid+i] = dp[mid+i-1]
		}
	}
	fmt.Println("debug:", dp)
	ans := make([]byte, dp[slen-1])
	j := 0
	count := len(ans)
	for i := slen - 1; i > 0 && count > 0; i-- {
		if dp[i] > dp[i-1] {
			ans[j] = s[i]
			j++
			count -= 2
			continue
		}
		if count == 1 {
			ans[j] = s[i]
			count -= 1
		}
	}
	for i := len(ans) - 1; i >= len(ans)/2; i-- {
		ans[i] = ans[len(ans)-i-1]
	}
	return string(ans)
}
