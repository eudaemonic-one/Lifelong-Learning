# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> List[List[int]]:
        ans = []
        def dfs(root, sum, path):
            if not root.left and not root.right and root.val == sum:
                ans.append(path)
                return
            sum -= root.val
            if root.left:
                dfs(root.left, sum, path+[root.left.val])
            if root.right:
                dfs(root.right, sum, path+[root.right.val])
        if not root:
            return []
        dfs(root, sum, [root.val])
        return ans

    def pathSum(self, root: TreeNode, total: int) -> List[List[int]]:
        if not root:
            return []
        stack, ans = [(root, [root.val])], []
        while stack:
            node, path = stack.pop()
            if not node.left and not node.right and sum(path) == total:
                ans.append(path)
            if node.left:
                stack.append((node.left, path+[node.left.val]))
            if node.right:
                stack.append((node.right, path+[node.right.val]))
        return ans
