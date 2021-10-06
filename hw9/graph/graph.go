package graph

import (
	"strconv"
	"strings"
)

type Graph struct {
	IsDirected bool
	nodes      []Node
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
	g.IsDirected = true
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
		s = append(s, "directed true")
	} else {
		s = append(s, "directed false")
	}

	for _, node := range g.nodes {
		for _, attr := range node.attrs {
			s = append(s, attr.name, string(rune(attr.val)))
		}
		for _, edge := range node.Incident {
			for _, attr := range edge.attrs {
				s = append(s, attr.name, "value", string(rune(attr.val)))
			}
		}
	}
	return s
}
