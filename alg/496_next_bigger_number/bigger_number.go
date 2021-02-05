package bigger_number

/*
给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

*/
/*思路：
使用栈和map来解决，首先遍历nums2的所有值，尝试压入栈
当待push元素A比栈顶元素大时，弹出栈顶元素，并记录栈顶元素的bigger值为元素A
继续比较栈顶元素，直到元素A小于栈顶元素或者栈为空，将元素A push
A比栈顶元素小时，将A push

优化点：不需要记录所有元素的bigger值，只需要记录在nums1的bigger值，减少空间占用

思考：什么场景下需要用单调栈来考虑解决呢？

*/
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	s := NewStack()
	m := constructNums1Map(nums1)
	for _, num := range nums2 {
		for {
			if s.IsEmpty() {
				s.Push(num)
				break
			}
			top, _ := s.Peek()
			if top.(int) < num {
				if _, ok := m[top.(int)]; ok {
					m[top.(int)] = num
				}
				s.Pop()
			} else {
				s.Push(num)
				break
			}
		}
	}
	return deconstructNums1Map(nums1, m)
}
func deconstructNums1Map(nums1 []int, m map[int]int) []int {
	res := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		res = append(res, m[v])
	}
	return res
}

/*首先构建map，对应的bigger值首先用-1来代替*/
func constructNums1Map(nums1 []int) map[int]int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num] = -1
	}
	return m
}
