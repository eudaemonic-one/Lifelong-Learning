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
        level_start = root
        while level_start:
            curr = level_start
            while curr:
                if curr.left:
                    curr.left.next = curr.right
                if curr.right and curr.next:
                    curr.right.next = curr.next.left
                curr = curr.next
            level_start = level_start.left
        return root
