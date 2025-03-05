package main

import "fmt"

// Nó da Trie que armazena os filhos e indica se é o fim de uma palavra
type TrieNode struct {
	children map[rune]*TrieNode // Mapeia cada caractere para seu nó filho
	isEnd    bool               // Indica se o nó marca o final de uma palavra
}

// Estrutura da Trie com a raiz da árvore
type Trie struct {
	root *TrieNode // Nó raiz da Trie
}

// Cria e retorna um novo nó da Trie
func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode), // Inicializa o mapa de filhos
		isEnd:    false,                    // Por padrão, não é fim de palavra
	}
}

// Cria e retorna uma nova Trie com a raiz inicializada
func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

// Insere uma palavra na Trie, criando nós conforme necessário
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word { // Itera por cada caractere da palavra
		// Se o caractere não existir como filho, cria um novo nó
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = newTrieNode()
		}
		// Avança para o nó filho correspondente
		node = node.children[ch]
	}
	// Marca o nó final como o fim de uma palavra
	node.isEnd = true
}

// Procura uma palavra na Trie; retorna true se a palavra foi inserida
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word { // Percorre a Trie buscando cada caractere
		if _, exists := node.children[ch]; !exists {
			return false // Se faltar um caractere, a palavra não existe
		}
		node = node.children[ch]
	}
	return node.isEnd // Retorna true apenas se o nó final marcar o fim de uma palavra
}

func main() {
	trie := NewTrie()     // Cria uma nova Trie
	trie.Insert("golang") // Insere palavras na Trie
	trie.Insert("estrutura")
	trie.Insert("dados")

	// Imprime o resultado das buscas na Trie
	fmt.Println("Busca 'golang':", trie.Search("golang")) // true
	fmt.Println("Busca 'gola':", trie.Search("gola"))     // false
	fmt.Println("Busca 'dados':", trie.Search("dados"))   // true
}
