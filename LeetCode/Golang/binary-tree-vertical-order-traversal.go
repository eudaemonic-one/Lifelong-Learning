/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    l, r := 0, 0
    dfs(root, 0, &l, &r)
    l = -l
    res := make([][]int, l+r+1)
    qnode := []*TreeNode{root}
    qcol := []int{0}
    for len(qnode) > 0 {
        node := qnode[0]
        qnode = qnode[1:]
        col := qcol[0]
        qcol = qcol[1:]
        res[col+l] = append(res[col+l], node.Val)
        if node.Left != nil {
            qnode = append(qnode, node.Left)
            qcol = append(qcol, col-1)
        }
        if node.Right != nil {
            qnode = append(qnode, node.Right)
            qcol = append(qcol, col+1)
        }
    }
    return res
}

func dfs(root *TreeNode, col int, l, r *int) {
    if root == nil {
        return
    }
    if col < *l {
        *l = col
    } else if col > *r {
        *r = col
    }
    dfs(root.Left, col-1, l, r)
    dfs(root.Right, col+1, l, r)
}
