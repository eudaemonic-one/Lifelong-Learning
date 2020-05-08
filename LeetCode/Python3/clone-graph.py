"""
# Definition for a Node.
class Node:
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors
"""
class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if not node:
            return None
        copy = Node(node.val, [])
        hashmap = {node: copy}
        stack = [node]
        while stack:
            node = stack.pop()
            for nei in node.neighbors:
                if nei not in hashmap:
                    tmp = Node(nei.val, [])
                    hashmap[nei] = tmp
                    hashmap[node].neighbors.append(tmp)
                    stack.append(nei)
                else:
                    hashmap[node].neighbors.append(hashmap[nei])
        return copy
