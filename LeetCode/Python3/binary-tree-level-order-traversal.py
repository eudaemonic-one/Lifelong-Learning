# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return None
        cnt1, cnt2 = 1, 0
        queue = [root]
        level = []
        ans = []
        while queue:
            node = queue.pop(0)
            if node.left:
                queue.append(node.left)
                cnt2 += 1
            if node.right:
                queue.append(node.right)
                cnt2 += 1
            level.append(node.val)
            if cnt1 > 0:
                cnt1 -= 1
            if cnt1 == 0:
                ans.append(level)
                level = []
                cnt1 = cnt2
                cnt2 = 0
        return ans

    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return None
        queue = [(0, root)]
        level = []
        ans = []
        while queue:
            depth, node = queue.pop(0)
            if node:
                if len(ans) == depth:
                    ans.append([])
                ans[depth].append(node.val)
                if node.left:
                    queue.append((depth+1, node.left))
                if node.right:
                    queue.append((depth+1, node.right))
        return ans
