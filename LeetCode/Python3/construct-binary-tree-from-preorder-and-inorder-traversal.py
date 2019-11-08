# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if inorder:
            root = TreeNode(preorder.pop(0))
            idx = inorder.index(root.val)
            root.left = self.buildTree(preorder, inorder[:idx])
            root.right = self.buildTree(preorder, inorder[idx+1:])
            return root
