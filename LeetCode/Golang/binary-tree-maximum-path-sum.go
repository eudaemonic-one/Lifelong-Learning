/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    var res = math.MinInt32
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	*res = max(left+right+root.Val, *res)
	return max(0, max(left, right)+root.Val)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}