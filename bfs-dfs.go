package main

import "fmt"

type Node struct {
	val   int
	edges []*Node
}

func BFS(graph *[]Node) {
	visited := make(map[int]bool, len(*graph))
	queue := make([]*Node, 0, len(*graph))
	queue = append(queue, &(*graph)[0])
	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]
		// marking as visited
		if visited[currentNode.val] {
			continue
		}
		visited[currentNode.val] = true
		fmt.Println("Visited node", currentNode.val)
		// visiting Node = appending its edges to the "queue"
		for _, node := range currentNode.edges {
			queue = append(queue, node)
		}
	}
}
func DFS(currentNode *Node, visited *map[int]bool) {
	if (*visited)[currentNode.val] {
		return
	}
	(*visited)[currentNode.val] = true
	fmt.Println("Visited node", currentNode.val)
	for _, node := range currentNode.edges {
		DFS(node, visited)
	}
}

func main() {
	nodes := make([]Node, 6, 6)
	// 0 connected to 3,4
	nodes[0].edges = append(nodes[0].edges, &nodes[3], &nodes[4])
	// 1 connected to 2
	nodes[1].edges = append(nodes[1].edges, &nodes[1])
	// 2 connected to 4
	nodes[2].edges = append(nodes[2].edges, &nodes[4], &nodes[1])
	// 3 connected to 5
	nodes[3].edges = append(nodes[3].edges, &nodes[5], &nodes[0])
	// and now the reversals
	nodes[4].edges = append(nodes[4].edges, &nodes[0], &nodes[2])
	nodes[2].edges = append(nodes[2].edges, &nodes[1])
	nodes[5].edges = append(nodes[5].edges, &nodes[3])

	for i := 0; i < 6; i++ {
		nodes[i].val = i
	}
	fmt.Println("BFS")
	BFS(&nodes)
	fmt.Println("DFS")
	visited := make(map[int]bool, len(nodes))
	DFS(&nodes[0], &visited)

	// Testing Graph
	coolGraph := CreateGraph()
	coolGraph.AddNode(0)
	coolGraph.AddNode(1)
	coolGraph.AddNode(2)
	coolGraph.AddNode(3)
	coolGraph.AddNode(4)
	coolGraph.AddNode(5)
	coolGraph.AddEdge(0, 4)
	coolGraph.AddEdge(0, 5)
	coolGraph.AddEdge(1, 2)
	coolGraph.AddEdge(2, 4)
	coolGraph.AddEdge(3, 5)
	fmt.Println("Node 0's neighbors:", coolGraph.nodes[0].edges)
	fmt.Println("Node 1's neighbors:", coolGraph.nodes[1].edges)
	fmt.Println("Node 3's neighbors:", coolGraph.nodes[3].edges)
}

// I should probably build a way to define & work with graph in an easier way
type (
	Graph     struct{ nodes map[int]*GraphNode }
	GraphNode struct {
		val   int
		edges []*GraphNode
		// optionally I could go for referecing outNeighbors and inNeighbors; let's prioritize using less memory now
	}
)

func CreateGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*GraphNode),
	}
}
func (graph *Graph) AddNode(nodeId int) {
	_, exists := graph.nodes[nodeId]
	if exists {
		fmt.Println("Cannot add, already exists")
		return
	}
	newNode := &GraphNode{val: nodeId, edges: make([]*GraphNode, 0, 5)}
	graph.nodes[nodeId] = newNode
	fmt.Println("New node added")
}
func (graph *Graph) AddEdge(nodeId1 int, nodeId2 int) {
	node1, exists1 := graph.nodes[nodeId1]
	node2, exists2 := graph.nodes[nodeId2]
	if !exists1 || !exists2 {
		fmt.Println("Cannot add edge when a node doesn't exists")
		return
	}
	node1.edges = append(node1.edges, node2)
	node2.edges = append(node2.edges, node1)

}
func (node *GraphNode) removeEdge(neighbor *GraphNode) {
	// instead of linear searching, if we have the guarantee of adding / keeping nodes in order, we could do binary search
	for i, edge := range node.edges {
		if edge == neighbor {
			node.edges = append(node.edges[:i], node.edges[i+1:]...)
			return
		}
	}
	fmt.Println("Could not remove edge")
}
func (graph *Graph) RemoveEdge(node *GraphNode, neighbor *GraphNode) {
	// first find the node
	if _, exists := graph.nodes[node.val]; !exists {
		fmt.Println("Node does not exist")
		return
	}
	node.removeEdge(neighbor)
	neighbor.removeEdge(node)
}
