class Node:
    def __init__(self, val: int, key: int):
        self.val = val
        self.key = key
        self.neighbors = []
        return None

    def add_neighbor(self, neighbor):
        self.neighbors.append(neighbor)
