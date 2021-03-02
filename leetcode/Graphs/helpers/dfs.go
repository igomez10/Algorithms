package main

import "container/list"

func DFS(root *TreeNode) []int {
	toVisit := list.New()

	toVisit.PushBack(root)
	answer := []int{}
	for toVisit.Len() > 0 {
		current := toVisit.Remove(toVisit.Back()).(*TreeNode)
		if current == nil {
			continue
		}
		answer = append(answer, current.Val)
		if current.Right != nil {
			toVisit.PushBack(current.Right)
		}
		if current.Left != nil {
			toVisit.PushBack(current.Left)
		}
	}

	return answer
}
