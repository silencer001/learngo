package leetcode

import "math"

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*思考：
合并有序列表，考虑构建一个小顶堆，每次弹出堆顶元素*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func heapify(h []*ListNode, index int) {
	son := index*2 + 1
	for son < len(h) {
		if son+1 < len(h) && h[son].Val > h[son+1].Val {
			son++
		}
		if h[index].Val > h[son].Val {
			h[index], h[son] = h[son], h[index]
			index = son
			son = son*2 + 1
		} else {
			break
		}
	}
}
func min_heapify(h []*ListNode) {
	for i := len(h) / 2; i >= 0; i-- {
		heapify(h, i)
	}
}
func mergeKLists(lists []*ListNode) *ListNode {
	fake := new(ListNode)
	p := fake
	heap := make([]*ListNode, 0)
	for _, v := range lists {
		if v != nil {
			heap = append(heap, v)
		}
	}
	if len(heap) == 0 {
		return nil
	}
	min_heapify(heap)
	//弹出堆顶元素，如果
	for len(heap) > 1 {
		p.Next = heap[0]
		p = p.Next
		if heap[0].Next != nil {
			heap[0] = heap[0].Next
		} else {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		}
		heapify(heap, 0) //对堆顶重新初始化
	}
	p.Next = heap[0].Next
	return fake.Next
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tmp := make([]int, len(nums1)+len(nums2))
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums1[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}
	if i < len(nums1) {
		for i < len(nums1) {
			tmp[k] = nums1[i]
			k++
			i++
		}
	}
	if j < len(nums2) {
		for j < len(nums2) {
			tmp[k] = nums2[j]
			k++
			j++
		}
	}
	mid1 := len(tmp) / 2
	mid2 := (len(tmp) - 1) / 2
	return (float64(tmp[mid1]) + float64(tmp[mid2])) / 2
}

func divide(dividend int, divisor int) int {
	divid := int64(dividend)
	divis := int64(divisor)
	var flag bool = true
	if (divid < 0 && divis > 0) || (divid > 0 && divis < 0) {
		flag = false
	}
	if divid < 0 {
		divid = -divid
	}
	if divid > math.MaxInt32 {
		divid = math.MaxInt32
	}
	if divisor < 0 {
		divisor = -divisor
	}

	result := 0
	for divid > divis {
		tmp := divis
		tmpres := 1
		for {
			if divid > (tmp << 1) {
				tmpres = tmpres << 1
				tmp = tmp << 1
			} else {
				break
			}
		}
		result += tmpres
		divid -= tmp
	}
	if flag {
		return result
	} else {
		return -result
	}
}
