package graphs

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Value     int
	Neighbors []*Node
}

func InsertNode(graph *Graph, node *Node) {
	graph.Nodes = append(graph.Nodes, node)
}

func AddEdge(node1 *Node, node2 *Node) {
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}

func DeleteEdge(node1 *Node, node2 *Node) {
	node1.Neighbors = RemoveNode(node1.Neighbors, node2)
	node2.Neighbors = RemoveNode(node2.Neighbors, node1)
}

func RemoveNode(nodes []*Node, node *Node) []*Node {
	for i := range nodes {
		if nodes[i] == node {
			return append(nodes[:i], nodes[i+1:]...)
		}
	}
	return nodes
}
