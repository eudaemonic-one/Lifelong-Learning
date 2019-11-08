# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def recoverTree(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        stack = []
        prev = TreeNode(-sys.maxsize)
        first = second = None
        while True:
            while root:
                stack.append(root)
                root = root.left
            if not stack:
                break
            node = stack.pop()
            if prev and prev.val >= node.val:
                if not first:
                    first = prev
                second = node
            prev = node
            root = node.right
        first.val, second.val = second.val, first.val
