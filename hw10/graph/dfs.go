package graph

import (
	"strconv"
	"strings"
)

type Graph struct {
	IsDirected bool
	nodes      []Node
	Visited    map[string]bool
}

type Node struct {
	ID       string
	attrs    []Attr
	Incident []edge
}

type edge struct {
	attrs          []Attr
	Source, Target *Node
}

type Attr struct {
	name string
	val  int
}

func New(s string) *Graph {
	var lines = strings.Split(s, "\n")
	var g Graph
	for _, v := range lines {
		words := strings.Split(v, " ")
		if len(words) == 4 {
			if words[3] == "true" {
				g.IsDirected = true
			}
		} else if len(words) == 5 {
			var e []edge
			x, _ := strconv.Atoi(words[2])
			a1 := Attr{words[1], x}
			y, _ := strconv.Atoi(words[4])
			a2 := Attr{words[3], y}
			n := Node{words[0], []Attr{a1, a2}, e}
			g.nodes = append(g.nodes, n)
		} else if len(words) == 3 {
			for np, n := range g.nodes {
				if n.ID == string(words[0][0]) {
					for mp, m := range g.nodes {
						if m.ID == string(words[0][1]) {
							e1, _ := strconv.Atoi(words[2])
							ege := edge{[]Attr{{words[0], e1}}, &g.nodes[np], &g.nodes[mp]}
							g.nodes[np].Incident = append(g.nodes[np].Incident, ege)
						}
					}

				}
			}
		}
	}
	return &g
}

func (g Graph) NumEdges() int {
	var count int
	for _, node := range g.nodes {
		count = +len(node.Incident)
	}
	return count
}

func (g Graph) HasEdge(n1, n2 string) bool {
	tmp := false
	for _, node := range g.nodes {
		for _, name := range node.Incident {
			if name.Source.ID == n1 && name.Target.ID == n2 {
				tmp = true
			}
		}
	}
	return tmp
}

func (g Graph) String() []string {
	var s []string
	if g.IsDirected == true {
		s = append(s, "directed true\n")
	} else {
		s = append(s, "directed false\n")
	}
	for _, node := range g.nodes {
		for _, attr := range node.attrs {
			s = append(s, "f", attr.name, string(rune(attr.val)), "\n")
		}
		for _, edge := range node.Incident {
			for _, attr := range edge.attrs {
				s = append(s, attr.name, "value", string(rune(attr.val)), "\n")
			}
		}
	}
	return s
}

func (g *Graph) ClearEdges() {
	for cnt1 := 0; cnt1 < len(g.nodes); cnt1++ {
		g.nodes[cnt1].Incident = nil
	}
}

func (g *Graph) Node(name string) *Node {
	var n *Node
	for _, node := range g.nodes {
		if node.ID == name {
			*n = node
		}
	}
	return n
}

func (g *Graph) Edge(ID0, ID1 string) *edge {
	var e *edge
	for _, node := range g.nodes {
		for _, edge := range node.Incident {
			if edge.attrs[0].name == ID0+ID1 {
				*e = edge
			}
		}
	}
	return e
}

func (n *Node) String() []string {
	var s []string
	s = append(s, n.ID, "attrs:", n.attrs[0].name, string(rune(n.attrs[0].val)), n.attrs[1].name, string(rune(n.attrs[1].val)), "incident:")
	for _, edge := range n.Incident {
		s = append(s, edge.Target.ID+"(value "+string(rune(edge.attrs[0].val))+")")
	}
	return s
}

func (g Graph) NbrsOf(name string) []Node {
	var nodes []Node
	for _, node := range g.nodes {
		if node.ID == name {
			for _, edge := range node.Incident {
				if edge.Source.attrs[0].name == name {
					nodes = append(nodes, *edge.Target)
				} else if edge.Target.attrs[0].name == name {
					nodes = append(nodes, *edge.Source)
				}
			}
		}
	}
	return nodes
}

func explore(g Graph, name string) []string {
	var s []string
	s = append(s, "("+name)
	g.Visited[name] = true
	for _, node := range g.nodes {
		for _, edge := range node.Incident {
			if g.Visited[edge.attrs[0].name] == false {
				explore(g, edge.Target.ID)
			}
		}
	}
	return s
}

func (g *Graph) ClearVisited() {
	for _, node := range g.nodes {
		g.Visited[node.ID] = false
	}
}

func DFS(g Graph) []string {
	var s []string
	var s1 []string
	for _, node := range g.nodes {
		g.Visited[node.ID] = false
	}
	for _, node := range g.nodes {
		if g.Visited[node.ID] == false {
			s1 = explore(g, node.ID)
			for _, word := range s1 {
				s = append(s, word)
			}
		}
	}
	return s
}
