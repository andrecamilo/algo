package main

import "fmt"

// Busca em Profundidade (DFS)
// Implementação simples de DFS utilizando recursão.

func dfsUtil(grafo map[int][]int, visitados map[int]bool, no int) {
	visitados[no] = true
	fmt.Printf("%d ", no)
	for _, vizinho := range grafo[no] {
		if !visitados[vizinho] {
			dfsUtil(grafo, visitados, vizinho)
		}
	}
}

func dfs(grafo map[int][]int, inicio int) {
	visitados := make(map[int]bool)
	dfsUtil(grafo, visitados, inicio)
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

	fmt.Print("DFS começando no nó 1: ")
	dfs(grafo, 1)
	fmt.Println()
}
