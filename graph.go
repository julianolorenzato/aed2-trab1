package main

import "fmt"

const MAX = 20
const INFINITY = 2 << 64

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

func (g *Graph) Dijkstra(src, dest int) {
	isVisited := [MAX]bool{false}
	path := [MAX]int{-1}
	costs := [MAX]int{INFINITY}
	costs[src] = 0

	for i := 0; i < MAX; i++ {
		if isVisited[i] {
			continue
		}

		for i, e := range g.Vertices[i].Edges {
			if e == nil {
				continue
			}
		}
	}
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

func main() {
	g := NewGraph()

	var option int

	for option != 4 {
		fmt.Println("1 - Adicionar vértice")
		fmt.Println("2 - Adicionar aresta")
		fmt.Println("3 - Mostrar valores")
		fmt.Println("4 - Sair")
		fmt.Println()

		fmt.Scanf("%d", &option)
		fmt.Println()

		switch option {
		case 1:
			var value int
			fmt.Println("Digite o valor do vértice")
			fmt.Scanf("%d", &value)

			if value < 0 {
				fmt.Println("O valor do vértice deve ser positivo")
			} else {
				err := g.AddVertex(value)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			fmt.Println()

		case 2:
			var srcIndex, destIndex, weight int
			fmt.Println("Digite o índice do vértice de origem")
			fmt.Scanf("%d", &srcIndex)
			fmt.Println("Digite o índice do vértice de destino")
			fmt.Scanf("%d", &destIndex)
			fmt.Println("Digite o peso da aresta")
			fmt.Scanf("%d", &weight)
			fmt.Println()

			err := g.AddEdge(srcIndex, destIndex, weight)
			if err != nil {
				fmt.Println(err.Error())
			}

		case 3:
			g.PrintAll()
			fmt.Println()
		}
	}
}
