import collections
print()


def printMatrix(matrix):
    for row in matrix:
        print(row)


def returnShortestPath(initialMatrix, origin, target):
    matrix = initialMatrix
    matrix[origin[0]][origin[1]] = 1
    matrix[target[0]][target[1]] = 1
    return matrix


def getShortestDistanceToAllPoints(initialMatrix, origin):
    matrix = buildMatrix(len(initialMatrix[0]), len(initialMatrix), 99)
    queue = collections.deque()
    queue.append((origin[0], origin[1]))
    matrix[origin[0]][origin[1]] = 10
    while len(queue) > 0:
        curr = queue.popleft()
        ci = curr[0]
        cj = curr[1]

        if ci > 0 and matrix[ci-1][cj] > matrix[ci][cj] + 1 and initialMatrix[ci-1][cj] >= 0:
            matrix[ci-1][cj] = matrix[ci][cj] + 1
            queue.append((ci-1, cj))

        if cj > 0 and matrix[ci][cj-1] > matrix[ci][cj] + 1 and initialMatrix[ci][cj-1] >= 0:
            matrix[ci][cj-1] = matrix[ci][cj] + 1
            queue.append((ci, cj-1))

        if ci < len(matrix)-1 and matrix[ci+1][cj] > matrix[ci][cj] + 1 and initialMatrix[ci+1][cj] >= 0:
            matrix[ci+1][cj] = matrix[ci][cj] + 1
            queue.append((ci+1, cj))

        if cj < len(matrix)-1 and matrix[ci][cj+1] > matrix[ci][cj] + 1 and initialMatrix[ci][cj+1] >= 0:
            matrix[ci][cj+1] = matrix[ci][cj] + 1
            queue.append((ci, cj+1))

    return matrix


def buildMatrix(x, y, startValue):
    finalList = []
    for i in range(y):
        row = []
        for j in range(x):
            row.extend([startValue])
        finalList.append(row)
    return finalList


def getShortestPath(matrix, origin, target):

    i, j = origin[0], origin[1]
    finalI, finalJ = target[0], target[1]

    nextPosition = (i, j)
    while i != finalI or j != finalJ:
        smallestCost = 2**31
        matrix[i][j] = 0
        if i > 0:
            if matrix[i-1][j] < smallestCost and matrix[i-1][j] != 0:
                nextPosition = (i-1, j)
                smallestCost = matrix[i-1][j]

        if j > 0:
            if matrix[i][j-1] < smallestCost and matrix[i][j-1] != 0:
                nextPosition = (i, j-1)
                smallestCost = matrix[i][j-1]

        if i < len(matrix)-1:
            if matrix[i+1][j] < smallestCost and matrix[i+1][j] != 0:
                nextPosition = (i+1, j)
                smallestCost = matrix[i+1][j]

        if j < len(matrix[i])-1:
            if matrix[i][j+1] < smallestCost and matrix[i][j+1] != 0:
                nextPosition = (i, j+1)
                smallestCost = matrix[i][j+1]

        if nextPosition is None:
            break

        i = nextPosition[0]
        j = nextPosition[1]

    return matrix


matrix = buildMatrix(10, 10, 0)

matrix[4][5] = -1
matrix[5][5] = -1
matrix[3][4] = -1
matrix[4][6] = -1
matrix[4][7] = -1
matrix[4][8] = -1
matrix[5][8] = -1
matrix[5][9] = -1
matrix[2][3] = -1
matrix[4][3] = -1
matrix[6][5] = -1

adjustedMatrix = getShortestDistanceToAllPoints(matrix, (0, 0))

shortestPath = getShortestPath(adjustedMatrix, (len(
    adjustedMatrix)-1, len(adjustedMatrix)-1), (0, 0))

printMatrix(shortestPath)

# shortestPath = returnShortestPath(matrix, (0, 0), (9, 9))

# printMatrix(shortestPath)
