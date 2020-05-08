/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
    res := 0
    dfs(root, nil, 0, &res)
    return res
}

func dfs(root *TreeNode, parent *TreeNode, path int, res *int) {
    if root == nil {
        return
    }
    if parent != nil && root.Val == parent.Val + 1 {
        path++
    } else {
        path = 1
    }
    *res = max(*res, path)
    dfs(root.Left, root, path, res)
    dfs(root.Right, root, path, res)
    return
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
