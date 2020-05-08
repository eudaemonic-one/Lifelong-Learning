# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False
        if not root.left and not root.right and root.val == sum:
            return True
        sum -= root.val
        return self.hasPathSum(root.left, sum) or self.hasPathSum(root.right, sum)

    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False
        stack = [(root, sum)]
        path = 0
        while stack:
            node, path = stack.pop()
            if not node.left and not node.right and node.val == path:
                return True
            if node.left:
                stack.append((node.left, path-node.val))
            if node.right:
                stack.append((node.right, path-node.val))
        return False
