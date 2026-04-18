package dijkstra

type Edge struct {
	NodeA  *GraphNode
	NodeB  *GraphNode
	Weight int
}

type GraphNode struct {
	Key       int
	Neighbors map[int]*Edge
}

// option 1, represent graph as matrix nxm where matrix[n][m] shows the weight to
// go from n to m. the problem is that if we scale this to lots of node, we will
// grow space n*m basically n**2

// option 2. represent graph as a map of keys as strings, and values as node
// data structures. each data structure will have a key with the node identifier
// and

func BuildGraphFromEdges(edges [][3]int) map[int]*GraphNode {
	newGraph := map[int]*GraphNode{}

	for i := range edges {
		origin := edges[i][0]
		target := edges[i][1]
		weight := edges[i][2]

		// if origin node doesnt exist, create it
		if _, exists := newGraph[origin]; !exists {
			newGraph[origin] = &GraphNode{
				Key:       origin,
				Neighbors: map[int]*Edge{},
			}
		}
		if _, exists := newGraph[target]; !exists {
			newGraph[target] = &GraphNode{
				Key:       target,
				Neighbors: map[int]*Edge{},
			}
		}

		originNode := newGraph[origin]
		targetNode := newGraph[target]

		// create originToTargetEdge
		edgeOriginToTarget := &Edge{
			NodeA:  originNode,
			NodeB:  targetNode,
			Weight: weight,
		}
		edgeTargetToOrigin := &Edge{
			NodeA:  targetNode,
			NodeB:  originNode,
			Weight: weight,
		}

		// add edges to nodes
		originNode.Neighbors[target] = edgeOriginToTarget
		targetNode.Neighbors[origin] = edgeTargetToOrigin
	}

	return newGraph
}

func FindShortestPathBFS(graph map[int]*GraphNode, startKey, endKey int) []int {
	// add initial node to the queue
	// visit every node using bfs and keep track of the distance
	type Node struct {
		Val      *GraphNode
		Path     []int
		SumSoFar int
		Next     *Node
	}
	origin := graph[startKey]
	head := &Node{
		Val:      origin,
		Path:     []int{startKey},
		SumSoFar: 0,
	}
	tail := head
	seen := map[int]bool{
		origin.Key: true,
	}

	for head != nil {
		current := head

		if current.Val.Key == endKey {
			return current.Path
		}

		// add all neighbors to end of queue
		for i := range current.Val.Neighbors {
			currentEdge := current.Val.Neighbors[i]
			currentTarget := currentEdge.NodeB.Key
			if seen[currentTarget] == false {
				seen[currentTarget] = true
				newNode := &Node{
					Val:  currentEdge.NodeB,
					Path: append(current.Path, currentEdge.NodeB.Key),
				}
				tail.Next = newNode
				tail = tail.Next
			}
		}

		head = head.Next
	}

	return []int{-1}
}
