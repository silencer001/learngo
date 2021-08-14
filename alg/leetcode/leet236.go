package leetcode

/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ppath := postOrder(root, p)
	qpath := postOrder(root, q)
	i := 1
	for ; i < len(ppath) && i < len(qpath); i++ {
		if ppath[i] != qpath[i] {
			return ppath[i-1]
		}
	}
	if len(ppath) < len(qpath) {
		return ppath[len(ppath)-1]
	} else {
		return qpath[len(qpath)-1]
	}
}

/*先写个后续遍历的循环方法吧*/

func postOrder(root *TreeNode, target *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	p := root
	var prev *TreeNode
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			if p.Val == target.Val {
				return stack
			}
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if p.Right == nil || p.Right == prev { //p没有右子树或者已经遍历过p的右子树了，此时遍历p
				//fmt.Println(p.Val)
				prev = p
				p = nil
			} else {
				stack = append(stack, p) //重新push后，遍历右子树
				p = p.Right
			}
		}
	}
	return nil
}
