package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Definição de uma aresta
type Aresta struct {
	alvo int
	peso int
}

// Representação do grafo usando lista de adjacências
type Grafo map[int][]Aresta

// Item para a fila de prioridade
type Item struct {
	no, distancia int
}

// Fila de prioridade implementada com heap
type FilaPrioridade []Item

func (pq FilaPrioridade) Len() int { return len(pq) }
func (pq FilaPrioridade) Less(i, j int) bool {
	return pq[i].distancia < pq[j].distancia
}
func (pq FilaPrioridade) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *FilaPrioridade) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *FilaPrioridade) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(grafo Grafo, inicio int) map[int]int {
	distancias := make(map[int]int)
	for no := range grafo {
		distancias[no] = math.MaxInt64
	}
	distancias[inicio] = 0

	pq := &FilaPrioridade{}
	heap.Init(pq)
	heap.Push(pq, Item{no: inicio, distancia: 0})

	for pq.Len() > 0 {
		atual := heap.Pop(pq).(Item)
		noAtual := atual.no
		distAtual := atual.distancia

		if distAtual > distancias[noAtual] {
			continue
		}

		for _, aresta := range grafo[noAtual] {
			novoDist := distAtual + aresta.peso
			if novoDist < distancias[aresta.alvo] {
				distancias[aresta.alvo] = novoDist
				heap.Push(pq, Item{no: aresta.alvo, distancia: novoDist})
			}
		}
	}

	return distancias
}

func main() {
	grafo := Grafo{
		1: {{alvo: 2, peso: 2}, {alvo: 3, peso: 4}},
		2: {{alvo: 3, peso: 1}, {alvo: 4, peso: 7}},
		3: {{alvo: 5, peso: 3}},
		4: {{alvo: 6, peso: 1}},
		5: {{alvo: 4, peso: 2}, {alvo: 6, peso: 5}},
		6: {},
	}

	inicio := 1
	distancias := dijkstra(grafo, inicio)
	fmt.Printf("Distâncias mínimas a partir do nó %d:\n", inicio)
	for no, dist := range distancias {
		fmt.Printf("Nó %d: %d\n", no, dist)
	}
}
