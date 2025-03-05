package main

import (
	"container/list"
	"fmt"
)

// Busca em Largura (BFS)
// Implementação simples de BFS utilizando fila.

func bfs(grafo map[int][]int, inicio int) {
	visitados := make(map[int]bool)
	fila := list.New()

	visitados[inicio] = true
	fila.PushBack(inicio)

	for fila.Len() > 0 {
		elemento := fila.Front()
		no := elemento.Value.(int)
		fmt.Printf("%d ", no)
		fila.Remove(elemento)

		for _, vizinho := range grafo[no] {
			if !visitados[vizinho] {
				visitados[vizinho] = true
				fila.PushBack(vizinho)
			}
		}
	}
}

func main() {
	grafo := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {6},
		6: {},
	}

	fmt.Print("BFS começando no nó 1: ")
	bfs(grafo, 1)
	fmt.Println()
}
