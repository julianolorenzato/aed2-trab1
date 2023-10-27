package main

import (
	"fmt"

	graph "github.com/julianolorenzato/aed2-trab1"
)

func main() {
	g := graph.NewGraph()

	var option int

	for {
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

			err := g.AddVertex(value)
			if err != nil {
				fmt.Println(err.Error())
			}

		case 2:
			var srcIndex, destIndex, weight int
			fmt.Println("Digite o índice do vértice de origem")
			fmt.Scanf("%d", &srcIndex)
			fmt.Println("Digite o índice do vértice de destino")
			fmt.Scanf("%d", &destIndex)
			fmt.Println("Digite o peso da aresta")
			fmt.Scanf("%d", &weight)

			err := g.AddEdge(srcIndex, destIndex, weight)
			if err != nil {
				fmt.Println(err.Error())
			}

		case 3:
			g.PrintAll()

		default:
			break
		}
	}
}
