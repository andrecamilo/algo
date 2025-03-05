package main

import (
	"container/heap"
	"fmt"
)

/*
Algoritmo de Huffman
O algoritmo de Huffman é utilizado para compressão de dados,
construindo uma árvore onde os símbolos mais frequentes têm códigos binários menores.
O exemplo a seguir constrói a árvore com base em uma frequência simulada para os caracteres de um texto.
*/

type Node struct {
	char        rune
	frequency   int
	left, right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

func buildHuffmanTree(frequencies map[rune]int) *Node {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for ch, freq := range frequencies {
		heap.Push(pq, &Node{char: ch, frequency: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(pq).(*Node)
		right := heap.Pop(pq).(*Node)
		merged := &Node{
			char:      0, // nó interno
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}
		heap.Push(pq, merged)
	}
	return heap.Pop(pq).(*Node)
}

func generateHuffmanCodes(root *Node, code string, codes map[rune]string) {
	if root == nil {
		return
	}
	// Se é folha, associa o caractere ao código
	if root.left == nil && root.right == nil {
		codes[root.char] = code
		return
	}
	generateHuffmanCodes(root.left, code+"0", codes)
	generateHuffmanCodes(root.right, code+"1", codes)
}

func main() {
	texto := "this is an example for huffman encoding"
	frequency := make(map[rune]int)
	for _, ch := range texto {
		frequency[ch]++
	}

	root := buildHuffmanTree(frequency)
	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	fmt.Println("Códigos de Huffman:")
	for ch, code := range codes {
		fmt.Printf("'%c': %s\n", ch, code)
	}
}
