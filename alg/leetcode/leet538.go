package leetcode

/*
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var help func(root *TreeNode, cur int) int
	help = func(root *TreeNode, cur int) int {
		if root == nil {
			return cur
		}
		right := help(root.Right, cur)
		root.Val += right
		cur = root.Val
		return help(root.Left, root.Val)
	}
	help(root, 0)
	return root
}
