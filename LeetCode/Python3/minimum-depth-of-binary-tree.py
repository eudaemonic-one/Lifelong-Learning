# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def minDepth(self, root: TreeNode) -> int:
        def dfs(root):
            if not root:
                return 0
            l = dfs(root.left)
            r = dfs(root.right)
            if l == 0 or r == 0:
                return max(l, r) + 1
            else:
                return min(l, r) + 1
        return dfs(root)

    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        queue = collections.deque([root])
        depth = 1
        while queue:
            for _ in range(len(queue)):
                node = queue.popleft()
                if not node.left and not node.right:
                    return depth
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            depth += 1
        return -1
