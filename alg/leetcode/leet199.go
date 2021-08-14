package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	qu := make([]*TreeNode, 0)
	var ans []int
	qu = append(qu, root)
	for len(qu) > 0 {
		ans = append(ans, qu[len(qu)-1].Val)
		tmp := make([]*TreeNode, 0)
		for len(qu) > 0 {
			cur := qu[0]
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
			qu = qu[1:]
		}
		qu = tmp
	}
	return ans
}
