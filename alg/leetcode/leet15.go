package leetcode

import "sort"

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		//去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second, third := first+1, len(nums)-1; second < third; {
			//first、second去重的话，third绝对不会重复，所以无需判断third是否和前一个third重复
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			res := nums[first] + nums[second] + nums[third]
			if res > 0 {
				third--
			} else if res < 0 {
				second++
			} else {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				third--
				second++
			}
		}
	}
	return ans
}
