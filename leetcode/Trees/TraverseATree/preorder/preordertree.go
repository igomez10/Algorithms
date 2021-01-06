package main

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
	head := &NodeToVisit{Node: root}
	tail := &NodeToVisit{}
	head.Next = tail

	ans := []int{}

	for head != nil && head.Node != nil && head != tail {
		front := head
		head = head.Next

		ans = append(ans, front.Node.Val)
		if front.Node.Right != nil {
			oldHead := head
			head = &NodeToVisit{Node: front.Node.Right, Next: oldHead}
		}

		if front.Node.Left != nil {
			oldHead := head
			head = &NodeToVisit{Node: front.Node.Left, Next: oldHead}
		}
	}
	return ans
}
