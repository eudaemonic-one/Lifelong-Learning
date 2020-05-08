/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    if (root == nil) {
        return 0
    }
    var path, sum = 0, 0
    nodeStack := make([]*TreeNode, 0)
    nodeStack = append(nodeStack, root)
    pathStack := make([]int, 0)
    pathStack = append(pathStack, root.Val)
    for len(nodeStack) > 0 {
        root = nodeStack[len(nodeStack)-1]
        nodeStack = nodeStack[0:len(nodeStack)-1]
        path = pathStack[len(pathStack)-1]
        pathStack = pathStack[0:len(pathStack)-1]
        if (root.Left == nil && root.Right == nil) {
            sum += path
        }
        if (root.Left != nil) {
            nodeStack =append(nodeStack, root.Left)
            pathStack = append(pathStack, path*10 + root.Left.Val)
        }
        if (root.Right != nil) {
            nodeStack =append(nodeStack, root.Right)
            pathStack = append(pathStack, path*10 + root.Right.Val)
        }
    }
    return sum
}
