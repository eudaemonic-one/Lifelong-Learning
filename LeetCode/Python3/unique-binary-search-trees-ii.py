# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        def dfs(start, end):
            if start >= end:
                return [None]
            ans = []
            for i in range(start, end): # root node
                for l in dfs(start, i): # left child tree
                    for r in dfs(i+1, end): # right child tree
                        root = TreeNode(i)
                        root.left = l
                        root.right = r
                        ans.append(root)
            return ans
        if n == 0:
            return None
        return dfs(1, n+1)
