from collections import deque


class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


def build_sample_tree() -> Node:
    root = Node(0)
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)
    node5 = Node(5)
    node6 = Node(6)

    root.left = node1
    root.right = node2
    node1.left = node3
    node1.right = node4
    node2.left = node5
    node2.right = node6

    return root


def inorder_traversal(root: Node) -> list[int]:
    res = []
    queue = deque()
    queue.append(root)

    while queue:
        current = queue.popleft()
        res.append(current.value)
        if current.left:
            queue.append(current.left)
        if current.right:
            queue.append(current.right)

    return res


def preorder_traversal(root: Node) -> list[int]:
    res = []
    queue = deque()
    queue.append(root)
    while queue:
        current = queue.pop()
        res.append(current.value)
        if current.right:  # append right first
            queue.append(current.right)
        if current.left:  # append left after (so we can pop it first)
            queue.append(current.left)

    return res


def postorder_traversal(root: Node) -> list[int]:
    res = []
    queue = deque()
    queue.append(root)
    stack = deque()
    while queue:
        current = queue.pop()
        stack.append(current)

        if current.left:
            queue.append(current.left)
        if current.right:
            queue.append(current.right)

    while stack:
        current = stack.pop()
        res.append(current.value)

    return res


def test_inorder():
    tree = build_sample_tree()
    assert inorder_traversal(tree) == [0, 1, 2, 3, 4, 5, 6]


def test_preorder():
    tree = build_sample_tree()
    assert preorder_traversal(root=tree) == [0, 1, 3, 4, 2, 5, 6]


def test_postorder():
    tree = build_sample_tree()
    assert postorder_traversal(root=tree) == [3, 4, 1, 5, 6, 2, 0]
