package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Estrutura do Bloom Filter
type BloomFilter struct {
	bitset []bool
	k      int // número de funções hash
	m      int // tamanho do bitset
}

// Cria um novo Bloom Filter
func NewBloomFilter(m, k int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, m),
		k:      k,
		m:      m,
	}
}

// Função hash baseada em SHA256 com um "seed" variável
func (bf *BloomFilter) hash(data []byte, seed int) int {
	h := sha256.New()
	h.Write([]byte{byte(seed)})
	h.Write(data)
	sum := h.Sum(nil)
	// Converte os primeiros 4 bytes para um inteiro e aplica módulo m
	return int(binary.BigEndian.Uint32(sum) % uint32(bf.m))
}

// Adiciona um item (como string) ao Bloom Filter
func (bf *BloomFilter) Add(item string) {
	data := []byte(item)
	for i := 0; i < bf.k; i++ {
		pos := bf.hash(data, i)
		bf.bitset[pos] = true
	}
}

// Verifica se o item pode estar presente
func (bf *BloomFilter) Contains(item string) bool {
	data := []byte(item)
	for i := 0; i < bf.k; i++ {
		pos := bf.hash(data, i)
		if !bf.bitset[pos] {
			return false
		}
	}
	return true
}

func main() {
	bf := NewBloomFilter(1000, 3) // 1000 bits, 3 funções hash
	bf.Add("golang")
	bf.Add("estruturas de dados")

	fmt.Println("Contém 'golang'?", bf.Contains("golang"))
	fmt.Println("Contém 'python'?", bf.Contains("python"))
}
