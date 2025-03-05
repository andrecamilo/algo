package main

import "fmt"

func buscaBinaria(arr []int, alvo int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == alvo {
			return mid
		} else if arr[mid] < alvo {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	alvo := 7
	indice := buscaBinaria(arr, alvo)
	if indice != -1 {
		fmt.Printf("Elemento %d encontrado na posição %d\n", alvo, indice)
	} else {
		fmt.Printf("Elemento %d não encontrado\n", alvo)
	}
}
