package graph

import (
	"os"
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

func (g *Graph) init(file string) {
	f, _ := os.ReadFile(file)
	f1 := string(f)
	var A, B []Attr
	var src, trgt *Node
	var E []edge
	g.IsDirected = true
	lines := strings.Split(f1, "\n")
	for _, v := range lines {
		words := strings.Split(v, " ")
		if len(words) == 2 {
			if words[1] == "true" {
				g.IsDirected = true
			} else {
				g.IsDirected = false
			}
		} else if len(words) == 5 {
			av, _ := strconv.Atoi(words[2])
			a := Attr{name: words[1], val: av}
			A = append(A, a)
			bv, _ := strconv.Atoi(words[4])
			b := Attr{name: words[3], val: bv}
			A = append(A, b)
			n := Node{ID: words[0], attrs: A}
			g.nodes = append(g.nodes, n)
		} else if len(words) == 3 {
			av, _ := strconv.Atoi(words[2])
			a := Attr{name: words[0], val: av}
			B = append(B, a)
			letter := strings.Split(words[0], "")
			for _, val := range g.nodes {
				if val.ID == letter[0] {
					src = &val
				} else if val.ID == letter[1] {
					trgt = &val
				}
			}
			e := edge{attrs: B, Source: src, Target: trgt}
			for _, nds := range g.nodes {
				if e.Source.ID == nds.ID {
					E = append(E, e)
					nds.Incident = E
				}
			}
		}
	}
}
func (g Graph) NumEdges() int {
	return len(g.nodes)
}

func (g Graph) HasEdge(a string, b string) bool {
	var tmp bool
	name := a + b
	for _, v := range g.nodes {
		for _, v1 := range v.Incident {
			for _, v2 := range v1.attrs {
				if v2.name == name {
					tmp = true
				} else {
					tmp = false
				}
			}
		}
	}
	return tmp
}
