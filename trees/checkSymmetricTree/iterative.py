# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
import collections


class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:

        if root is None:
            return True
        if root.left is None and root.right is None:
            return True

        queue = collections.deque()
        queue.append(root)
        queue.append(root)

        leftNode = 0
        rightNode = 0

        while len(queue) > 0:
            leftNode = queue.pop()
            rightNode = queue.pop()

            # compare key
            if leftNode.val != rightNode.val:
                return False

            if leftNode.left is not None and rightNode.right is not None:
                queue.append(leftNode.left)
                queue.append(rightNode.right)
            elif leftNode.left is not None or rightNode.right is not None:
                return False

            if leftNode.right is not None and rightNode.left is not None:
                queue.append(leftNode.right)
                queue.append(rightNode.left)
            elif leftNode.right is not None or rightNode.left is not None:
                return False

        return True
