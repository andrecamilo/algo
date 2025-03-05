package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Estrutura do Bloom Filter, uma estrutura probabilística para verificar pertencimento de itens
type BloomFilter struct {
	bitset []bool // Vetor de bits que representa os bits definidos
	k      int    // Número de funções hash utilizadas
	m      int    // Tamanho do vetor de bits (bitset)
}

// Cria um novo Bloom Filter com m bits e k funções hash
func NewBloomFilter(m, k int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, m), // Inicializa o vetor de bits com falso
		k:      k,
		m:      m,
	}
}

// Função hash que utiliza SHA256 e uma semente para gerar índices
func (bf *BloomFilter) hash(data []byte, seed int) int {
	h := sha256.New()
	h.Write([]byte{byte(seed)}) // Incorpora a semente
	h.Write(data)               // Incorpora os dados
	sum := h.Sum(nil)
	// Converte os primeiros 4 bytes para um inteiro e utiliza o módulo para restringir ao tamanho do bitset
	return int(binary.BigEndian.Uint32(sum) % uint32(bf.m))
}

// Adiciona um item (como string) ao Bloom Filter, definindo os bits correspondentes
func (bf *BloomFilter) Add(item string) {
	data := []byte(item)
	for i := 0; i < bf.k; i++ {
		pos := bf.hash(data, i)
		bf.bitset[pos] = true
	}
}

// Verifica se um item está presente no Bloom Filter; pode haver falsos positivos
func (bf *BloomFilter) Contains(item string) bool {
	data := []byte(item)
	for i := 0; i < bf.k; i++ {
		pos := bf.hash(data, i)
		if !bf.bitset[pos] {
			return false // Se algum dos bits não estiver definido, o item definitivamente não está presente
		}
	}
	return true // Se todos os bits estiverem definidos, o item pode estar presente (mas pode ser um falso positivo)
}

func main() {
	bf := NewBloomFilter(1000, 3) // Cria um Bloom Filter com 1000 bits e 3 funções hash
	bf.Add("golang")              // Adiciona itens ao filtro
	bf.Add("estruturas de dados")

	// Verifica a presença dos itens
	fmt.Println("Contém 'golang'?", bf.Contains("golang")) // true
	fmt.Println("Contém 'python'?", bf.Contains("python")) // Provavelmente false
}
