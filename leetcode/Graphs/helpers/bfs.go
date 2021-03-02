package main

import "container/list"

func BFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	answers := []int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(*TreeNode)
		if current == nil {
			continue
		}

		answers = append(answers, current.Val)
		if current.Left != nil {
			queue.PushBack(current.Left)
		}
		if current.Right != nil {
			queue.PushBack(current.Right)
		}
	}
	return answers
}
