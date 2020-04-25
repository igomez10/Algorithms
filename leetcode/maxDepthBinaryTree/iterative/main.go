package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	if root == nil {
		return maxDepth
	}

	queue := list.New()
	seen := map[*TreeNode]int{}

	queue.PushBack(root)
	seen[root] = 1

	for queue.Len() > 0 {
		first := queue.Back()
		fmt.Println("Removing", queue.Len())
		curr := queue.Remove(first).(*TreeNode)
		fmt.Println("Removed", queue.Len())
		// fmt.Println(queue.Len())
		// fmt.Printf("%+v\n", queue.Front())

		depth := seen[curr]
		if depth > maxDepth {
			maxDepth = depth
		}

		if curr.Left != nil {
			queue.PushBack(curr.Left)
			seen[curr.Left] = depth + 1
		}

		if curr.Right != nil {
			queue.PushBack(curr.Right)
			seen[curr.Right] = depth + 1
		}
	}

	return maxDepth
}

func main() {
	n0 := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}

	n0.Right = n1
	n1.Right = n2
	n2.Right = n3
	n3.Right = n4
	n4.Right = n5

	maxDepth(n0)

}
