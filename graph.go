package graph

import (
	"fmt"
)

const MAX = 20

type Graph struct {
	Vertices [MAX]*vertex // Array de valores do tipo *vertex
}

type vertex struct {
	Val   any        // O valor do vertice pode ser de qualquer tipo
	Edges [MAX]*edge // Array de valores do tipo *edge
}

type edge struct {
	Weight     int
	DestVertex *vertex
}

// Construtor
func NewGraph() *Graph {
	return &Graph{
		Vertices: [20]*vertex{nil},
	}
}

func (g *Graph) AddVertex(val any) error {
	newVertex := &vertex{Val: val, Edges: [20]*edge{nil}}

	for i, v := range g.Vertices {
		if v == nil {
			g.Vertices[i] = newVertex
			return nil
		}
	}

	return fmt.Errorf("O limite máximo de %v vértices já foi atingido.", MAX)
}

func (g *Graph) AddEdge(srcIndex, destIndex, weight int) error {
	switch {
	case srcIndex < 0 || srcIndex > MAX-1:
		return fmt.Errorf("O vertice origem deve estar entre 0 e %v.", MAX-1)

	case destIndex < 0 || destIndex > MAX-1:
		return fmt.Errorf("O vertice destino deve estar entre 0 e %v.", MAX-1)

	case g.Vertices[srcIndex] == nil:
		return fmt.Errorf("O vertice %v não existe.", srcIndex)

	case g.Vertices[destIndex] == nil:
		return fmt.Errorf("O vertice %v não existe.", destIndex)

	case g.Vertices[srcIndex].Edges[destIndex] != nil:
		return fmt.Errorf("A aresta que liga o vertice %v ao vertice %v já existe.", srcIndex, destIndex)
	}

	newEdge := &edge{
		Weight:     weight,
		DestVertex: g.Vertices[destIndex],
	}

	g.Vertices[srcIndex].Edges[destIndex] = newEdge

	return nil
}

func (g *Graph) PrintAll() {
	for i, v := range g.Vertices {
		if v == nil {
			continue
		}

		fmt.Printf("Vértice %v (%v) Arestas:", i, v.Val)

		for i, e := range v.Edges {
			if e == nil {
				continue
			}

			fmt.Printf(" [Destino: %v Peso: %v]", i, e.Weight)
		}

		fmt.Println()
	}
}
