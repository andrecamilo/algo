package main

import "fmt"

/*
Union-Find (Disjoint Set Union)
O algoritmo de Union-Find (ou Disjoint Set Union – DSU) é utilizado para manter o controle de conjuntos disjuntos e é amplamente usado em problemas de conectividade (por exemplo, para detectar ciclos em grafos).
*/

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

func main() {
	uf := NewUnionFind(10)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(4, 5)

	fmt.Println("Find(3):", uf.Find(3)) // Deve pertencer ao mesmo conjunto de 1,2,3
	fmt.Println("Find(5):", uf.Find(5)) // Conjunto de 4,5
	fmt.Println("Find(6):", uf.Find(6)) // 6 está isolado
}
