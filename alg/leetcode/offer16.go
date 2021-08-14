package leetcode

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。



示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	type helpFunc func(x float64, n int) float64
	var help helpFunc
	help = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}
		if n == 0 {
			return 1
		}
		tmp := help(x, n/2)
		if n%2 == 1 {
			return tmp * tmp * x
		} else {
			return tmp * tmp
		}
	}
	return help(x, n)
}
