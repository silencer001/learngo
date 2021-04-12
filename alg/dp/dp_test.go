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
