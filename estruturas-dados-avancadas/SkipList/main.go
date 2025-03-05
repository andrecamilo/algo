package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 16 // Nível máximo permitido na Skip List
const p = 0.5       // Probabilidade para aumentar o nível de um novo nó

// Nó da Skip List que contém o valor e um slice de ponteiros para os nós de níveis superiores
type Node struct {
	value   int     // Valor armazenado no nó
	forward []*Node // Slice de ponteiros para os nós avançados em cada nível
}

// Estrutura da Skip List que possui um nó cabeçalho e informações do nível atual
type SkipList struct {
	header *Node // Nó cabeçalho especial
	level  int   // Nível máximo atual da lista
}

// Cria um novo nó com o valor dado e um slice de ponteiros com tamanho "level+1"
func NewNode(value, level int) *Node {
	return &Node{
		value:   value,
		forward: make([]*Node, level+1), // Níveis de 0 até level
	}
}

// Cria uma nova Skip List com um nó cabeçalho inicializado
func NewSkipList() *SkipList {
	return &SkipList{
		header: NewNode(0, maxLevel), // Header com valor 0 e nível máximo
		level:  0,
	}
}

// Gera aleatoriamente um nível para um novo nó, limitado a maxLevel
func randomLevel() int {
	level := 0
	// Enquanto o número aleatório for menor que p, aumenta o nível
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// Insere um novo valor na Skip List mantendo a ordenação
func (sl *SkipList) Insert(value int) {
	update := make([]*Node, maxLevel+1) // Array para armazenar os nós que precisarão ter seus ponteiros atualizados
	current := sl.header

	// Percorre os níveis da Skip List de cima para baixo para encontrar a posição de inserção
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].value < value {
			current = current.forward[i]
		}
		update[i] = current // Armazena o nó atual em cada nível
	}

	current = current.forward[0]
	// Se o valor não estiver presente, insere-o
	if current == nil || current.value != value {
		lvl := randomLevel() // Define o nível do novo nó
		if lvl > sl.level {
			// Se o novo nível for maior, atualiza todos os níveis intermediários com o header
			for i := sl.level + 1; i <= lvl; i++ {
				update[i] = sl.header
			}
			sl.level = lvl
		}
		newNode := NewNode(value, lvl) // Cria o novo nó
		for i := 0; i <= lvl; i++ {
			// Atualiza os ponteiros para inserir o nó no nível i
			newNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = newNode
		}
	}
}

// Busca um valor na Skip List; retorna true se encontrado
func (sl *SkipList) Search(value int) bool {
	current := sl.header
	// Percorre os níveis da lista de cima para baixo
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].value < value {
			current = current.forward[i]
		}
	}
	current = current.forward[0] // Avança para o nível 0
	// Retorna true se o nó encontrado contiver o valor exato
	return current != nil && current.value == value
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente aleatória
	sl := NewSkipList()              // Cria uma nova Skip List
	sl.Insert(5)
	sl.Insert(15)
	sl.Insert(25)
	sl.Insert(35)

	// Realiza testes de busca
	fmt.Println("Search 25:", sl.Search(25)) // true
	fmt.Println("Search 30:", sl.Search(30)) // false
}
