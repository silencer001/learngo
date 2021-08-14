package leetcode

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jian-sheng-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*最值问题，考虑用动态规划，存在最优子结构和重复解
动态方程：f(n) = max(f(n-1)*1, f(n-2)*2, ..., f(1)*n)
几个特殊的初始值
f(1)=0
f(2)=1
f(3)=2
f(4)=4，此时和方程式对不上了，这个是为啥呢？是不是应该只考虑max(f(n-1)*1, f(n-2)*2, ...f(n/2)*(n-n/2))呢
*/
func cuttingRope(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			if dp[i] < dp[j]*(i-j) {
				dp[i] = dp[j] * (i - j)
			}
		}
	}
	return dp[n]
}

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

func cuttingRope1(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	var res int64 = 1
	for n > 4 {
		n -= 3
		res = (res * 3) % 1000000007
	}
	if n == 4 {
		res = ((res * 2) % 1000000007 * 2) % 1000000007
	} else if n == 3 {
		res = (res * 3) % 1000000007
	} else if n == 2 {
		res = (res * 2) % 1000000007
	}
	return int(res)
}
