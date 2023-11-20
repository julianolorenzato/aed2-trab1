package main

import (
	"fmt"
	"math"
	"slices"
)

const MAX = 20
const INFINITY = math.MaxInt32

type Graph struct {
	Vertices [MAX]*Vertex // Array de valores do tipo *Vertex
}

type Vertex struct {
	Val   any        // O valor do vertice pode ser de qualquer tipo
	Edges [MAX]*Edge // Array de valores do tipo *edge
}

type Edge struct {
	Weight int
	//DestVertex *vertex
}

// Construtor
func NewGraph() *Graph {
	return &Graph{
		Vertices: [MAX]*Vertex{nil},
	}
}

func (g *Graph) AddVertex(val any) error {
	newVertex := &Vertex{Val: val, Edges: [MAX]*Edge{nil}}

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

	newEdge := &Edge{
		Weight: weight,
		//DestVertex: g.Vertices[destIndex],
	}

	g.Vertices[srcIndex].Edges[destIndex] = newEdge

	return nil
}

func (g *Graph) Dijkstra(src, dest int) (minCost int, path []int) {
	var isVisited [MAX]bool
	var costs [MAX]int
	var previous [MAX]int

	for i := 0; i < MAX; i++ {
		costs[i] = INFINITY
	}
	costs[src] = 0

	for i := 0; i < MAX; i++ {
		previous[i] = -1
	}

	curr := src
	for curr != -1 {
		isVisited[curr] = true

		for i, edge := range g.Vertices[curr].Edges {
			if edge == nil {
				continue
			}

			foundCost := costs[curr] + edge.Weight

			if foundCost < costs[i] {
				costs[i] = foundCost
				previous[i] = curr
			}
		}

		curr = findNextIndex(costs, isVisited)
	}

	return costs[dest], buildPath(src, dest, previous)
}

func findNextIndex(costs [MAX]int, isVisited [MAX]bool) int {
	value := INFINITY
	index := -1

	for i, v := range costs {
		if isVisited[i] {
			continue
		}

		if v < value {
			value = v
			index = i
		}
	}

	return index
}

func buildPath(source, destiny int, previous [MAX]int) []int {
	path := make([]int, 0)
	path = append(path, destiny)

	curr := previous[destiny]
	for curr != source {
		path = append(path, curr)
		curr = previous[curr]
	}

	path = append(path, source)

	slices.Reverse(path)

	return path
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

	for option != 5 {
		fmt.Println("1 - Adicionar vértice")
		fmt.Println("2 - Adicionar aresta")
		fmt.Println("3 - Mostrar grafo")
		fmt.Println("4 - Custo mínimo entre dois vértices")
		fmt.Println("5 - Sair")
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
		case 4:
			var source, destiny int
			fmt.Println("Digite o índice do vértice de saída")
			fmt.Scanf("%d", &source)
			fmt.Println("Digite o índice do vértice de chegada")
			fmt.Scanf("%d", &destiny)
			minCost, path := g.Dijkstra(source, destiny)
			fmt.Println("Custo mínimo:", minCost)
			fmt.Print("Caminho de custo mínimo:")
			for _, v := range path {
				fmt.Print(" -> ", v)
			}
			fmt.Println()
			fmt.Println()
		}
	}
}
