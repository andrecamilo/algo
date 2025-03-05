package main

import "fmt"

// Estrutura da Fenwick Tree, utilizada para cálculos dinâmicos de prefix sum
type FenwickTree struct {
	tree []int // Array interno da árvore (1-indexed)
	n    int   // Número de elementos
}

// Cria uma nova Fenwick Tree para n elementos (o índice 0 não é utilizado)
func newFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, n+1),
		n:    n,
	}
}

// Atualiza a Fenwick Tree adicionando "delta" ao índice i
func (ft *FenwickTree) update(i, delta int) {
	for i <= ft.n {
		ft.tree[i] += delta // Atualiza o valor na posição i
		i += i & (-i)       // Move para o próximo índice relevante
	}
}

// Retorna a soma dos elementos de 1 até o índice i
func (ft *FenwickTree) query(i int) int {
	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= i & (-i) // Move para trás para somar o prefixo
	}
	return sum
}

// Consulta a soma no intervalo [left, right] (índices base 1)
func (ft *FenwickTree) rangeQuery(left, right int) int {
	return ft.query(right) - ft.query(left-1)
}

func main() {
	// Exemplo de array 1-indexed; a posição 0 é ignorada
	arr := []int{0, 1, 7, 3, 5, 2}
	n := len(arr) - 1       // Número de elementos efetivos
	ft := newFenwickTree(n) // Cria a Fenwick Tree
	for i := 1; i <= n; i++ {
		ft.update(i, arr[i]) // Constrói a árvore com os valores do array
	}
	// Consulta a soma dos elementos nos índices 2 a 4: 7 + 3 + 5 = 15
	fmt.Println("Soma no intervalo de 2 a 4:", ft.rangeQuery(2, 4))
}
