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
	if root == nil {
		return []int{}
	}
	tail := &NodeToVisit{}
	head := &NodeToVisit{Node: root, Next: tail}

	ans := []int{}
	for head != nil {
		pop := head
		ans = append(ans, pop.Node.Val)
		head = head.Next

		if pop.Node.Left != nil {
			head = &NodeToVisit{Node: pop.Node.Left, Next: head}
		}
		if pop.Node.Right != nil {
			tail.Next = &NodeToVisit{Node: head.Node.Right}
			tail = tail.Next
		}

	}

	return ans
}
