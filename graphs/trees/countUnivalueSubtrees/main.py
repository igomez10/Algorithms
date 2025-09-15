class treeNode():
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

# bottom-up, first check the leaf nodes and count them,
# then go up, if both children are "True" and root.val is
# equal to both children's values if exist, then root node
# is uniValue suntree node.


class Solution:
    def isUnivalueSubtree(self, root):
        if root is None:
            return True

        l = self.isUnivalueSubtree(root.left)
        r = self.isUnivalueSubtree(root.right)

        if l and r and (root.left is None or root.left.val == root.val) and (root.right is None or root.right.val == root.val):
            self.counter += 1
            return True

        return False

    def countUnivalSubtrees(self, root: TreeNode) -> int:
        self.counter = 0
        self.isUnivalueSubtree(root)
        return self.counter


s = Solution()

demo = """
        5(0)
       /  \
      /    \
     /      \
    1(1)     5(2)
   /   \        \
  5(3)  5(4)    5(6)
"""

# EXAMPLE 1
node0 = treeNode(5)
node1 = treeNode(1)
node2 = treeNode(5)
node3 = treeNode(5)
node4 = treeNode(5)

node6 = treeNode(5)


node2.right = node6

node1.left = node3
node1.right = node4

node0.left = node1
node0.right = node2

expectedNumvalTrees = 4

numUnivalTrees = s.countUnivalSubtrees(node0)

if numUnivalTrees == expectedNumvalTrees:
    print("Working as expected")
else:
    print("Not working as expected", numUnivalTrees, expectedNumvalTrees)
