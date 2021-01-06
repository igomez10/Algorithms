package main

import (
	"container/list"
)

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeToVisit struct {
	Node *TreeNode
	Next *NodeToVisit
}

func preorderTraversal(root *TreeNode) []int {
	l := list.New()
	l.PushFront(root)
	ans := []int{}
	for l.Len() > 0 {
		pop := l.Remove(l.Front()).(*TreeNode)
		ans = append(ans, pop.Val)
		if pop.Left != nil {
			l.PushFront(pop.Left)
		}
		if pop.Right != nil {
			l.PushBack(pop.Right)
		}
	}
	return ans
}
