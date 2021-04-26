package dp

import (
	"fmt"
	"testing"
)

func TestFindMaxForm(t *testing.T) {
	//fmt.Println(myAtoi("123"))
	//fmt.Println(myAtoi("-123"))
	//fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
	strs := []string{"10", "0001", "111001", "1", "0"}
	fmt.Println(findMaxForm(strs, 5, 3))
}

func TestCountPalindromicSubsequences(t *testing.T) {
	fmt.Println(countPalindromicSubsequences("bccb"))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("aabab"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("a123aasgasdfgwerasdfasdf"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("cccbb"))
}

func TestIsMatch(t *testing.T) {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxsubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func TestClimbStairs(t *testing.T) {
	fmt.Println("n=1  :", climbStairs1(1))
	fmt.Println("n=3  :", climbStairs1(3))
	fmt.Println("n=10  :", climbStairs1(10))

}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println("(()", longestValidParentheses("(()"))
	fmt.Println(")))", longestValidParentheses(")))"))
	fmt.Println("((", longestValidParentheses("(("))
	fmt.Println("(())", longestValidParentheses("(())"))
	fmt.Println("(()))", longestValidParentheses("(()))"))
	fmt.Println("(())())", longestValidParentheses("()(())))"))
}

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func TestIsMatch1(t *testing.T) {
	fmt.Println(isMatch1("aa", "a*"))
	fmt.Println(isMatch1("aa", ""))
	fmt.Println(isMatch1("aab", "*"))
	fmt.Println(isMatch1("acdcb", "a*c?b"))
	fmt.Println(isMatch1("adceb", "*a*b"))
	fmt.Println(isMatch1("aa", "a"))
	fmt.Println(isMatch1("", ""))
	fmt.Println(isMatch1("a", ""))
	fmt.Println(isMatch1("abcabczzzde", "*abc???de*"))
}

func TestMinDifficulty(t *testing.T) {
	fmt.Println(minDifficulty([]int{11, 111, 22, 222, 33, 333, 44, 444}, 6))
}

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("abcde", "ebcde"))
	fmt.Println(minDistance("sea", "eat"))
}
