/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    res := root.Val
    cur := root
    diff := float64(int(^uint(0) >> 31))
    for cur != nil {
        if abs(float64(cur.Val)-target) < diff {
            diff = abs(float64(cur.Val)-target)
            res = cur.Val
        }
        if target == float64(cur.Val) {
            return cur.Val
        } else if target < float64(cur.Val) {
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }
    return res
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}
