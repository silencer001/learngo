package dp

import (
	"fmt"
)

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

/*使用优化的贪心算法，求每个字符的最长回文子串，以字符为起点，向左右推进，相同时迭代，不同时终止*/
func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		var init bool = true
		left, right := i, i
		for {
			if init && left >= 0 && s[left] == s[i] {
				left--
				continue
			}
			if init && right < len(s) && s[right] == s[i] {
				right++
				continue
			}
			init = false
			if left >= 0 && right < len(s) && s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		//fmt.Println(s[left+1 : right])
		if len(max) < right-left-1 {
			max = s[left+1 : right]
		}
	}
	return max
}
func longestPalindrome3(s string) string {
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

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。


示例 1：

输入：s = "aa" p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:

输入：s = "aa" p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：

输入：s = "ab" p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4：

输入：s = "aab" p = "c*a*b"
输出：true
解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5：

输入：s = "mississippi" p = "mis*is*p*."
输出：false


提示：

0 <= s.length <= 20
0 <= p.length <= 30
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*典型的动态规划题目
求转移方程：
f(si,pj)= {f(si-1,pj-1),s[i-1]==p[j-1]} || {p[j]==* && s[i] match p[j-1] && f(si-1,pj) || {p[j]==* && f[si,pj-2]}}

怎么转变为状态转移矩阵呢
用字符串s的长度作为阶段，p的长度作为范围
*/
func isCharMatch(c byte, p byte) bool {
	if c == p || p == '.' {
		return true
	} else {
		return false
	}
}
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' && dp[0][i-2] {
			dp[0][i] = true
		}
	}

	for i := 1; i <= len(s); i++ {
		ispat := false
		for j := 1; j <= len(p); j++ {

			ch := s[i-1]
			pch := p[j-1]
			if pch != '*' {
				if isCharMatch(ch, pch) {
					dp[i][j] = dp[i-1][j-1]
					ispat = true
				}
			} else {
				if dp[i][j-2] || (dp[i-1][j] && isCharMatch(ch, p[j-2])) {
					dp[i][j] = true
					ispat = true
				}
			}
		}
		if !ispat {
			return false
		}
	}
	return dp[len(s)][len(p)]
}

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [0]
输出：0
示例 4：

输入：nums = [-1]
输出：-1
示例 5：

输入：nums = [-100000]
输出：-100000


提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
O(n)时间复杂度的算法：遍历整个数组，计算连续子序列之和为cur，如果cur小于0，则子序列从当前下标开始，
因为cur为负的子序列只会让最大子序列和变小，必然不属于最大子序列的子集；但cur不管正负，都可能是max
比较max 和cur大小
*/
func maxsubArray(nums []int) int {
	max := nums[0]
	//start, end := 0, 0
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			//start = i
			cur = nums[i]
		} else {
			cur += nums[i]
		}

		if max < cur {
			max = cur
		}

	}
	return max
}

/*使用动态规划的解法：
子序列存在最优子结构、重复子结构问题，因此考虑动态规划
转移方程：第n个值的最大连续连续值为
f(n)= nums[n] 当f(n-1)<0时，nums[n]+f(n-1),当f(n-1)>0
转移数组，下标代表n，值为包含n的最大子序列和
同样是O(n)时间复杂度，但空间复杂度为O(n)
*/

func maxsubArray1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
存在最优子结构和重复子结果，考虑动态规划
f(n-1)种方法再上一个台阶
f(n-2)种方法再上两个台阶
因此f(n)=f(n-1)+f(n-2)
*/
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*继续优化，使用O(1)的空间*/
func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp0 := 1
	dp1 := 1
	dpn := 0
	for i := 2; i <= n; i++ {
		dpn = dp0 + dp1
		dp0 = dp1
		dp1 = dpn
	}
	return dpn
}

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。



示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0


提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
针对每个)算动态转移方程
f(n)= 2 //当n-1为(且n-2为(
f(n)= 2+f(n-2) //当n-1为(且n-2为)
f(n)= 2+f(n-1)//当n-1为)，且n-f(n-1)为(
f(n)=  f(n-1)//当n-1为),且n-f(n-1)为)
*/

func longestValidParentheses(s string) int {
	var max int = 0
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
		max = 2
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if s[i-2] == '(' {
				dp[i] = 2
			} else {
				dp[i] = dp[i-2] + 2
			}
		} else { //)
			if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2
				if i-1-dp[i-1]-1 >= 0 {
					dp[i] += dp[i-1-dp[i-1]-1]
				}
			} else {
				dp[i] = 0
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：



输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func trap(height []int) int {
	findtop := func(h []int) int {
		max := 0
		for i := 1; i < len(h); i++ {
			if h[i] > h[max] {
				max = i
			}
		}
		return max
	}
	w := 0
	type recallFunc func(int, int, bool)
	var recall recallFunc
	recall = func(start, end int, startIsMax bool) {
		if end-start <= 1 {
			return
		}
		if end-start == 2 { //只有三个值时，直接计算出w
			debug := 0
			if startIsMax {
				if height[start+1] < height[end] {
					w += height[end] - height[start+1]
					debug = height[end] - height[start+1]
				}
			} else {
				if height[start+1] < height[start] { //end为最大值，如果start大于中间值，则有雨水
					w += height[start] - height[start+1]
					debug = height[start] - height[start+1]
				}
			}
			fmt.Println("slice:", height[start:end+1], " w=", debug)
			return
		}
		if startIsMax {
			debug := 0
			max2 := findtop(height[start+1:end+1]) + start + 1
			for i := start + 1; i < max2; i++ {
				w += height[max2] - height[i]
				debug += height[max2] - height[i]
			}
			fmt.Println("slice:", height[start:max2+1], " w=", w)
			recall(max2, end, true)
		} else {
			max2 := findtop(height[start:end]) + start
			debug := 0
			for i := max2 + 1; i < end; i++ {
				w += height[max2] - height[i]
				debug += height[max2] - height[i]
			}
			fmt.Println("slice:", height[start:max2+1], " w=", debug)
			recall(start, max2, false)
		}
	}

	max := findtop(height)
	recall(0, max, false)
	recall(max, len(height)-1, true)
	return w
}

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输出: false
"abcabczzzde"
"*abc???de*"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
s的下标为i，p的下标为j
转移方程f(i,j)=f(i-1,j-1)&&{match(s[i],p[j]} || f(i-1,j)&&p[j]==*
转移矩阵需要考虑空字符串的情况

*/
func isMatch1(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	dp := make([][]bool, sl+1)
	for i, _ := range dp {
		dp[i] = make([]bool, pl+1)
	}
	dp[0][0] = true
	for j := 1; j <= pl; j++ {
		if dp[0][j-1] && p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if dp[i-1][j-1] {
				if s[i-1] == p[j-1] || p[j-1] == '*' || p[j-1] == '?' {
					dp[i][j] = true
					continue
				}
			}
			if p[j-1] == '*' {
				if dp[i-1][j] || dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[sl][pl]
}

/*
你需要制定一份 d 天的工作计划表。工作之间存在依赖，要想执行第 i 项工作，你必须完成全部 j 项工作（ 0 <= j < i）。

你每天 至少 需要完成一项任务。工作计划的总难度是这 d 天每一天的难度之和，而一天的工作难度是当天应该完成工作的最大难度。

给你一个整数数组 jobDifficulty 和一个整数 d，分别代表工作难度和需要计划的天数。第 i 项工作的难度是 jobDifficulty[i]。

返回整个工作计划的 最小难度 。如果无法制定工作计划，则返回 -1 。



示例 1：



输入：jobDifficulty = [6,5,4,3,2,1], d = 2
输出：7
解释：第一天，您可以完成前 5 项工作，总难度 = 6.
第二天，您可以完成最后一项工作，总难度 = 1.
计划表的难度 = 6 + 1 = 7
示例 2：

输入：jobDifficulty = [9,9,9], d = 4
输出：-1
解释：就算你每天完成一项工作，仍然有一天是空闲的，你无法制定一份能够满足既定工作时间的计划表。
示例 3：

输入：jobDifficulty = [1,1,1], d = 3
输出：3
解释：工作计划为每天一项工作，总难度为 3 。
示例 4：

输入：jobDifficulty = [7,1,7,1,7,1], d = 3
输出：15
示例 5：

输入：jobDifficulty = [11,111,22,222,33,333,44,444], d = 6
输出：843


提示：

1 <= jobDifficulty.length <= 300
0 <= jobDifficulty[i] <= 1000
1 <= d <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-difficulty-of-a-job-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
第n天从0到i的最小难度为dp(i,n)
dp(i,n)=min(dp(i-k,n-1)+cost(k,i)),n-1<=k<=i
状态转移矩阵
以n为阶，i为状态；初始值dp[0][0]=0,dp[]
*/
func minDifficulty(jobDifficulty []int, d int) int {
	jobs := len(jobDifficulty)
	if jobs < d {
		return -1
	}
	dp := make([][]int, d)
	for i, _ := range dp {
		dp[i] = make([]int, jobs)
	}
	maxs := make([][]int, jobs)
	for i, _ := range maxs {
		maxs[i] = make([]int, jobs)
	}
	for i := 0; i < jobs; i++ {
		for j := i; j < jobs; j++ {
			if i == j {
				maxs[i][j] = jobDifficulty[j]
			} else {
				if jobDifficulty[j] > maxs[i][j-1] {
					maxs[i][j] = jobDifficulty[j]
				} else {
					maxs[i][j] = maxs[i][j-1]
				}
			}
		}
	}
	fmt.Println(maxs)
	for i := 0; i < jobs; i++ {
		dp[0][i] = maxs[0][i]
	}
	for n := 1; n < d; n++ {
		for j := 1; j < jobs; j++ {
			for k := 1; k <= j-n+1; k++ {
				if dp[n][j] == 0 || dp[n][j] > dp[n-1][j-k]+maxs[j-k+1][j] {
					dp[n][j] = dp[n-1][j-k] + maxs[j-k+1][j]
				}
			}
		}
	}
	for i := 0; i < d; i++ {
		fmt.Println(dp[i])
	}
	return dp[d-1][jobs-1]
}

/*背包问题*/
func pack(w int, n []int, v []int) int {
	dp := make([]int, w)
	for i, _ := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < len(n); i++ {
		for j := w - n[i]; j >= 0; j-- {
			if dp[j] != -1 {
				dp[j+n[i]] = max(dp[j]+v[i], dp[j+n[i]], -1)
			}
		}
	}
	for j := w; j >= 0; j-- {
		if dp[j] != -1 {
			return dp[j]
		}
	}
	return -1
}

/*
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minDistance(word1 string, word2 string) int {
	min := func(a, b, c int) int {
		k := a
		if k > b {
			k = b
		}
		if k > c {
			k = c
		}
		return k
	}
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	dp := make([][]int, len(word1))
	for i, _ := range dp {
		dp[i] = make([]int, len2)
	}
	if word1[0] == word2[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < len1; i++ {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}
	for j := 1; j < len2; j++ {
		if word2[j] == word1[0] {
			dp[0][j] = j
		} else {
			dp[0][j] = dp[0][j-1] + 1
		}
	}
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			x := 0
			if word1[i] == word2[j] {
				x = dp[i-1][j-1]
			} else {
				x = dp[i-1][j-1] + 1
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, x)
		}
	}
	return dp[len1-1][len2-1]
}
