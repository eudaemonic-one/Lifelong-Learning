# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return None
        queue = collections.deque([root])
        direction = 1
        level = []
        ans = []
        while queue:
            for _ in range(len(queue)):
                node = queue.popleft()
                if node:
                    level.append(node.val)
                    if node.left:
                        queue.append(node.left)
                    if node.right:
                        queue.append(node.right)
            if level:
                ans.append(level[::direction])
                direction *= -1
                level.clear()
        return ans
