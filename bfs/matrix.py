print()


def printMatrix(matrix):
    for row in matrix:
        print(row)


def returnShortestPath(initialMatrix, origin, target):
    matrix = initialMatrix
    x = origin[0]
    y = origin[1]
    matrix[0][0] = 1
    matrix[target[0]][target[1]] = 1
    return matrix


def buildMatrix(x, y, startValue):
    finalList = []
    for i in range(y):
        row = []
        for j in range(x):
            row.extend([startValue])
        finalList.append(row)
    return finalList


matrix = buildMatrix(10, 10, 0)
shortestPath = returnShortestPath(matrix, (0, 0), (9, 9))

printMatrix(shortestPath)
