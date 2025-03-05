package main

import "fmt"

// Função para mesclar dois subarrays ordenados
func merge(esq, dir []int) []int {
	resultado := make([]int, 0, len(esq)+len(dir))
	i, j := 0, 0
	for i < len(esq) && j < len(dir) {
		if esq[i] < dir[j] {
			resultado = append(resultado, esq[i])
			i++
		} else {
			resultado = append(resultado, dir[j])
			j++
		}
	}
	resultado = append(resultado, esq[i:]...)
	resultado = append(resultado, dir[j:]...)
	return resultado
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	meio := len(arr) / 2
	esq := mergeSort(arr[:meio])
	dir := mergeSort(arr[meio:])
	return merge(esq, dir)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Vetor original:", arr)
	ordenado := mergeSort(arr)
	fmt.Println("Vetor ordenado:", ordenado)
}
