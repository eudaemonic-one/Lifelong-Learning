"""
# Definition for a Node.
class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return None
        queue = collections.deque([root])
        while queue:
            l = len(queue)
            for i in range(l):
                node = queue.popleft()
                if i < l-1:
                    node.next = queue[0]
                else:
                    node.next = None
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return root

    def connect(self, root: 'Node') -> 'Node':
        cur, next_level_cur, next_level_head = root, None, None
        while cur:
            if cur.left:
                if next_level_cur:
                    next_level_cur.next = cur.left
                else:
                    next_level_head = cur.left
                next_level_cur = cur.left
            if cur.right:
                if next_level_cur:
                    next_level_cur.next = cur.right
                else:
                    next_level_head = cur.right
                next_level_cur = cur.right
            if cur.next:
                cur = cur.next
            else:
                cur = next_level_head
                next_level_head = None
                next_level_cur = None
        return root
