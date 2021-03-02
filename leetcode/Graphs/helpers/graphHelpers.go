package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children map[int]bool
}

func buildTree(input []interface{}) *TreeNode {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	type buildState struct {
		node *TreeNode
		idx  int
	}

	states := list.New()
	root := &TreeNode{Val: input[0].(int)}
	states.PushBack(buildState{node: root, idx: 0})

	for states.Len() > 0 {
		current := states.Remove(states.Front()).(buildState)

		if current.idx*2+1 < len(input) && input[current.idx*2+1] != nil {
			current.node.Left = &TreeNode{Val: input[current.idx*2+1].(int)}
			states.PushBack(buildState{node: current.node.Left, idx: current.idx*2 + 1})
		}

		if current.idx*2+2 < len(input) && input[current.idx*2+2] != nil {
			current.node.Right = &TreeNode{Val: input[current.idx*2+2].(int)}
			states.PushBack(buildState{node: current.node.Right, idx: current.idx*2 + 2})
		}
	}
	return root
}

func buildGraph(numVertices int, edges [][]int) map[int]*Node {
	dictVertices := map[int]*Node{}
	for i := 1; i <= numVertices; i++ {
		dictVertices[i] = &Node{Val: i, Children: map[int]bool{}}
	}

	for _, e := range edges {
		origin := e[0]
		target := e[1]
		dictVertices[origin].Children[target] = true
	}
	return dictVertices
}

func pathExist(origin, target int, graph map[int]*Node, path map[int]bool) bool {
	path[origin] = true
	defer func() {
		path[origin] = false
	}()

	children := graph[origin].Children

	for child := range children {
		if child == target {
			return true
		}
		path[child] = true
		exist := pathExist(child, target, graph, path)
		if exist {
			return true
		}
		path[child] = false
	}
	return false
}

func detectCycleUtil(origin int, graph map[int]*Node, visited, stack []bool) bool {
	visited[origin] = true
	stack[origin] = true

	children := graph[origin].Children
	for child := range children {
		if stack[child] || detectCycleUtil(child, graph, visited, stack) {
			return true
		}
	}
	return false
}

func DetectCycle(numVertices int, graph map[int]*Node) bool {
	for origin := 1; origin <= numVertices; origin++ {
		visited := make([]bool, numVertices+1)
		stack := make([]bool, numVertices+1)
		if detectCycleUtil(origin, graph, visited, stack) {
			return true
		}
	}
	return false
}
