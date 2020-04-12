import collections
#  bfs for a binary search tree.


# how to build a binary search tree
# objects
# an array of items


class treeNode():
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


# lets create a demo tree with the following structure

demo = """
        7
       / \
      /   \
     /     \
    4       10
   / \     / \
  0   5   9   14
"""

# EXAMPLE 1
node0 = treeNode(0)
node5 = treeNode(5)
node9 = treeNode(9)
node14 = treeNode(14)

node4 = treeNode(4)
node4.left = node0
node4.right = node5

node10 = treeNode(10)
node10.left = node9
node10.right = node14

node7 = treeNode(7)
node7.left = node4
node7.right = node10


# EXAMPLE 2

# node1 = treeNode(1)
# node2 = treeNode(2)
# node3 = treeNode(3)
# node4 = treeNode(4)

# node4.left = node1
# node3.left = node2
# node3.right = node4


def printBFS(root):
    traverseBFS(root, printNodeValue)


def printDFS(root):
    traverseDFS(root, printNodeValue)


def printNodeValue(root):
    print(root.val)


def traverseBFS(root, action):
    queue = collections.deque()
    queue.append(root)

    while len(queue) > 0:
        curr = queue.popleft()
        if action is not None:
            action(curr)

        if curr.left is not None:
            queue.append(curr.left)

        if curr.right is not None:
            queue.append(curr.right)


def traverseDFS(root, action):
    queue = collections.deque()
    queue.append(root)

    while len(queue) > 0:
        curr = queue.pop()
        if action is not None:
            action(curr)

        if curr.left is not None:
            queue.append(curr.left)

        if curr.right is not None:
            queue.append(curr.right)


# bfs and add them if
#  maintain two queues left and right subtrees
#  main while loop where we continue if elements in left or right

def traversePostOrder(root):
    firstStack = collections.deque()
    secondStack = []

    firstStack.append(root)

    while len(firstStack) > 0:
        curr = firstStack.pop()
        secondStack.append(curr.val)

        if curr.left is not None:
            firstStack.append(curr.left)

        if curr.right is not None:
            firstStack.append(curr.right)

    return list(reversed(secondStack))


def getMaximumDepth(root):
    if root is None:
        return 0
    queue = collections.deque()
    queue.append((root, 0))
    maximumDepth = 0
    while len(queue) > 0:
        currNode, currDepth = queue.popleft()
        maximumDepth = max(maximumDepth, currDepth)

        if currNode.left is not None:
            queue.append((currNode.left, currDepth+1))

        if currNode.right is not None:
            queue.append((currNode.right, currDepth+1))

    return maximumDepth


def getValue(root):
    return root.val


# traverseBFS(node7, printNodeValue)
print()
# traverseDFS(node7, printNodeValue)
print()
print(traversePostOrder(node7))
print(getMaximumDepth(node7))
