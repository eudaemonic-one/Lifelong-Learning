# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sortedListToBST(self, head: ListNode) -> TreeNode:
        nums = []
        if not head:
            return None
        while head:
            nums.append(head.val)
            head = head.next
        def dfs(nums, l, r):
            if l < r:
                m = (l + r) // 2
                root = TreeNode(nums[m])
                root.left = dfs(nums, l, m)
                root.right = dfs(nums, m+1, r)
                return root
        return dfs(nums, 0, len(nums))
