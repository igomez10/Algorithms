package main

func numIslands(matrix [][]byte) int {
	counter := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				matrix = Paint(matrix, i, j)
				counter++
			}
		}
	}

	return counter
}

func Paint(matrix [][]byte, i, j int) [][]byte {
	matrix[i][j] = '0'
	if i > 0 && matrix[i-1][j] == '1' {
		matrix = Paint(matrix, i-1, j)
	}
	if i < len(matrix)-1 && matrix[i+1][j] == '1' {
		matrix = Paint(matrix, i+1, j)
	}
	if j > 0 && matrix[i][j-1] == '1' {
		matrix = Paint(matrix, i, j-1)
	}
	if j < len(matrix[i])-1 && matrix[i][j+1] == '1' {
		matrix = Paint(matrix, i, j+1)
	}

	return matrix
}
