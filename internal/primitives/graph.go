package primitives

import (
	"strings"
)

// Graph is the structure that contains nodes and edges
type Graph struct {
	nodes map[uint64]*Node
	edges map[Node][]*Node
}

// NewGraph returns a new empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[uint64]*Node),
		edges: make(map[Node][]*Node),
	}
}

// AddNode inserts a new node in the graph
func (g *Graph) AddNode(n *Node) *Node {
	g.nodes[n.Id] = n
	return n
}

// AddEdge inserts a new edge in the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

// String returns a string representation of the graph
func (g Graph) String() string {
	sb := strings.Builder{}
	for _, v := range g.nodes {
		sb.WriteString(v.String())
		sb.WriteString(" -> [ ")
		neighbors := g.edges[*v]
		for _, u := range neighbors {
			sb.WriteString(u.String())
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	return sb.String()
}
