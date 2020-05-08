# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return ""
        codec = ""
        queue = collections.deque()
        queue.append(root)
        while queue:
            node = queue.popleft()
            if not node:
                codec += "None,"
                continue
            codec += str(node.val) + ","
            queue.append(node.left)
            queue.append(node.right)
        return codec

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None
        nodes = data.split(",")
        root = TreeNode(int(nodes[0]))
        queue = collections.deque([])
        queue.append(root)
        i = 1
        while queue and i < len(nodes):
            node = queue.popleft()
            if nodes[i] != "None":
                left = TreeNode(int(nodes[i]))
                node.left = left
                queue.append(left)
            i += 1
            if i == len(nodes):
                return root
            if nodes[i] != "None":
                right = TreeNode(int(nodes[i]))
                node.right = right
                queue.append(right)
            i += 1
        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
