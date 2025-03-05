package main

import (
	"fmt"
)

// Definição de um nó em uma árvore AVL, que é uma árvore binária de busca balanceada
type AVLNode struct {
	key    int      // Valor armazenado no nó
	height int      // Altura do nó para cálculos de balanceamento
	left   *AVLNode // Ponteiro para o filho da esquerda
	right  *AVLNode // Ponteiro para o filho da direita
}

// Retorna a altura de um nó ou 0 se o nó for nil
func height(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

// Função que retorna o máximo entre dois inteiros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Realiza uma rotação à direita para balancear a árvore
func rightRotate(y *AVLNode) *AVLNode {
	x := y.left   // x será a nova raiz da subárvore
	T2 := x.right // T2 é a subárvore que será movida

	// Executa a rotação
	x.right = y
	y.left = T2

	// Atualiza as alturas dos nós afetados
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x // Retorna a nova raiz da subárvore
}

// Realiza uma rotação à esquerda para balancear a árvore
func leftRotate(x *AVLNode) *AVLNode {
	y := x.right // y será a nova raiz da subárvore
	T2 := y.left // T2 é a subárvore que será movida

	// Executa a rotação
	y.left = x
	x.right = T2

	// Atualiza as alturas dos nós afetados
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y // Retorna a nova raiz da subárvore
}

// Calcula o fator de balanceamento de um nó
func getBalance(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

// Insere uma chave na árvore AVL e realiza as rotações necessárias para manter o balanceamento
func insert(node *AVLNode, key int) *AVLNode {
	// Caso base: se o nó for nil, cria um novo nó
	if node == nil {
		return &AVLNode{key: key, height: 1}
	}
	if key < node.key {
		node.left = insert(node.left, key) // Insere na subárvore esquerda
	} else if key > node.key {
		node.right = insert(node.right, key) // Insere na subárvore direita
	} else {
		return node // Chaves duplicadas não são permitidas
	}

	// Atualiza a altura do nó atual
	node.height = 1 + max(height(node.left), height(node.right))
	balance := getBalance(node) // Calcula o fator de balanceamento

	// Caso Left Left
	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	// Caso Right Right
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	// Caso Left Right: rotação dupla (primeiro à esquerda, depois à direita)
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Caso Right Left: rotação dupla (primeiro à direita, depois à esquerda)
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node // Retorna o nó inalterado (após inserção e balanceamento)
}

// Realiza uma travessia em ordem (in-order) para imprimir a árvore
func inOrder(root *AVLNode) {
	if root != nil {
		inOrder(root.left)
		fmt.Printf("%d ", root.key)
		inOrder(root.right)
	}
}

func main() {
	var root *AVLNode
	keys := []int{10, 20, 30, 40, 50, 25} // Valores a serem inseridos
	for _, key := range keys {
		root = insert(root, key)
	}

	// Imprime os valores da árvore em ordem crescente
	fmt.Println("In-order traversal da AVL Tree:")
	inOrder(root)
	fmt.Println()
}
