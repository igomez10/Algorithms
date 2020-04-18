
import (
	"container/list"
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	queue := list.New()
	queue.PushFront([]int{sr, sc})
	originalColor := image[sr][sc]
    visited := map[[]int]
	for queue.Len() > 0 {

		first := queue.Front()
		current := first.Value.([]int)

		image[current[0]][current[1]] = newColor
		

		queue.Remove(first)
        printMatrix(image)
        
		if current[0] > 0 && image[current[0]-1][current[1]] == originalColor {
			queue.PushBack([]int{current[0]-1, current[1]})
		}

		if current[0] < len(image)-1 && image[current[0]+1][current[1]] == originalColor {
			queue.PushBack([]int{current[0]+1, current[1]})
		}

		if current[1] > 0 && image[current[0]][current[1]-1] == originalColor {
			queue.PushBack([]int{current[0], current[1] - 1})
		}

		if current[1] < len(image[sr])-1 && image[current[0]][current[1]+1] == originalColor {
			queue.PushBack([]int{current[0], current[1] + 1})
		}

	}

	return image

}

func printMatrix(m [][]int) {
	for i := range m {
		fmt.Println(m[i])
	}
	fmt.Println()
}
