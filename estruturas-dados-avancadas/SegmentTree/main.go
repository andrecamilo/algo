package main

import "fmt"

// Estrutura da Segment Tree para realizar consultas de soma em intervalos
type SegmentTree struct {
	tree []int // Armazena os valores da árvore
	n    int   // Tamanho ajustado do array original (próxima potência de 2)
}

// Constrói uma nova Segment Tree a partir do array fornecido
func newSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	size := 1
	// Encontra a próxima potência de 2 maior ou igual a n
	for size < n {
		size *= 2
	}
	// A árvore terá tamanho 2 * size
	tree := make([]int, 2*size)
	st := &SegmentTree{tree: tree, n: size}
	st.build(arr, 1, 0, st.n-1)
	return st
}

// Função recursiva para construir a Segment Tree
// index: posição atual na árvore;
// [left, right]: intervalo correspondente no array original
func (st *SegmentTree) build(arr []int, index, left, right int) {
	if left == right {
		// Se o índice estiver no array original, atribui o valor; se não, permanece 0
		if left < len(arr) {
			st.tree[index] = arr[left]
		}
		return
	}
	mid := (left + right) / 2
	// Constrói recursivamente a subárvore esquerda e direita
	st.build(arr, 2*index, left, mid)
	st.build(arr, 2*index+1, mid+1, right)
	// Soma os valores dos nós filhos para formar o nó atual
	st.tree[index] = st.tree[2*index] + st.tree[2*index+1]
}

// Função recursiva que consulta a soma dos elementos no intervalo [ql, qr]
// index, left, right definem o intervalo atual na árvore
func (st *SegmentTree) query(index, left, right, ql, qr int) int {
	// Se o intervalo atual não intersecta com [ql, qr], retorna 0
	if ql > right || qr < left {
		return 0
	}
	// Se o intervalo atual for totalmente contido em [ql, qr], retorna o valor do nó
	if ql <= left && right <= qr {
		return st.tree[index]
	}
	mid := (left + right) / 2
	// Combina as consultas dos filhos esquerdo e direito
	return st.query(2*index, left, mid, ql, qr) +
		st.query(2*index+1, mid+1, right, ql, qr)
}

// Função de interface que realiza consulta no intervalo [ql, qr]
func (st *SegmentTree) Query(ql, qr int) int {
	return st.query(1, 0, st.n-1, ql, qr)
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	// Cria a Segment Tree a partir do array
	st := newSegmentTree(arr)
	// Consulta a soma dos elementos no intervalo de índices 1 a 3: 3 + 5 + 7 = 15
	fmt.Println("Soma entre índices 1 e 3:", st.Query(1, 3))
}
