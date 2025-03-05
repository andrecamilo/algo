package main

import "fmt"

func buscaLinear(arr []int, alvo int) int {
	for i, valor := range arr {
		if valor == alvo {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 7, 9, 11}
	alvo := 7
	indice := buscaLinear(arr, alvo)
	if indice != -1 {
		fmt.Printf("Elemento %d encontrado na posição %d\n", alvo, indice)
	} else {
		fmt.Printf("Elemento %d não encontrado\n", alvo)
	}
}
