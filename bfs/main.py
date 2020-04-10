import collections
from matrix import printMatrix


class Solution():

    def countIslands(self, arr):
        for i, row in enumerate(arr):
            for j, item in enumerate(row):
                if item == 1:
                    s.paintConnectedIslands(arr, i, j)

        counter = 0
        for row in arr:
            for item in row:
                if item == 1:
                    counter += 1

        return counter

    def paintConnectedIslands(self, arr, i, j):
        # validate matrix and coordinates
        if arr == [] or i > len(arr) or j > len(arr[0]) or arr[i][j] != 1:
            return

        queue = collections.deque()
        queue.append((i, j))
        seen = {(i, j): True}

        while len(queue) > 0:
            current = queue.pop()
            ci, cj = current[0], current[1]

            if ci > 0 and (ci-1, cj) not in seen and arr[ci-1][cj] == 1:
                queue.append((ci-1, cj))
                arr[ci-1][cj] = 0
                seen[(ci-1, cj)] = True

            if cj > 0 and (ci, cj-1) not in seen and arr[ci][cj-1] == 1:
                queue.append((ci, cj-1))
                arr[ci][cj-1] = 0
                seen[(ci, cj-1)] = True

            if ci < len(arr)-1 and (ci+1, cj) not in seen and arr[ci+1][cj] == 1:
                queue.append((ci+1, cj))
                arr[ci+1][cj] = 0
                seen[(ci+1, cj)] = True

            if cj < len(arr[ci])-1 and (ci, cj+1) not in seen and arr[ci][cj+1] == 1:
                queue.append((ci, cj + 1))
                arr[ci][cj+1] = 0
                seen[(ci, cj+1)] = True


arr = [
    [1, 1, 1, 0],
    [1, 1, 1, 0],
    [1, 0, 0, 0],
    [1, 0, 0, 0],
    [1, 0, 0, 0],
    [1, 0, 0, 0],
    [1, 0, 0, 0],
    [1, 0, 0, 0],
]

expectedResult = 2

s = Solution()

printMatrix(arr)
