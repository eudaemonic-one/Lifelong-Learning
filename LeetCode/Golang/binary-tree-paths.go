/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)
    dfs(root, "", &res)
    return res
}

func dfs(root *TreeNode, path string, res *[]string) {
    if root == nil {
        return
    }
    path += strconv.Itoa(root.Val)
    if root.Left == nil && root.Right == nil {
        *res = append(*res, path)
    }
    if root.Left != nil {
        dfs(root.Left, path+"->", res)
    }
    if root.Right != nil {
        dfs(root.Right, path+"->", res)
    }
}
