class LRUCache:

    def __init__(self, capacity: int):
        self.heap = []
        self.dict = {}  # {key:indexInHeap}
        self.count = 0
        heapq.heapify(self.heap)

    def get(self, key: int) -> int:
        self.count += 1
        if key in self.dict:
            self.dict[key][1] = self.count
            self.heap
            return self.dict[key][0]

        return -1

    def put(self, key: int, value: int) -> None:
        self.count += 1
        self.dict[key] = [value, self.count]


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
