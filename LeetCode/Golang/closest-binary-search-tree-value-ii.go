/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestKValues(root *TreeNode, target float64, k int) []int {
    predecessor := make([]int, 0)
    successor := make([]int, 0)
    res := make([]int, 0)
    inorder(root, target, &predecessor)
    reverseInorder(root, target, &successor)
    for k > 0 {
        if len(predecessor) == 0 {
            res = append(res, successor[len(successor)-1])
            successor = append([]int{}, successor[:len(successor)-1]...)
        } else if len(successor) == 0 {
            res = append(res, predecessor[len(predecessor)-1])
            predecessor = append([]int{}, predecessor[:len(predecessor)-1]...)
        } else if math.Abs(float64(predecessor[len(predecessor)-1]) - target) < math.Abs(float64(successor[len(successor)-1]) - target) {
            res = append(res, predecessor[len(predecessor)-1])
            predecessor = append([]int{}, predecessor[:len(predecessor)-1]...)
        } else {
            res = append(res, successor[len(successor)-1])
            successor = append([]int{}, successor[:len(successor)-1]...)
        }
        k--
    }
    return res
}

func inorder(root *TreeNode, target float64, stack *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, target, stack)
    if float64(root.Val) > target {
        return
    }
    *stack = append(*stack, root.Val)
    inorder(root.Right, target, stack)
}

func reverseInorder(root *TreeNode, target float64, stack *[]int) {
    if root == nil {
        return
    }
    reverseInorder(root.Right, target, stack)
    if float64(root.Val) <= target {
        return
    }
    *stack = append(*stack, root.Val)
    reverseInorder(root.Left, target, stack)
}
