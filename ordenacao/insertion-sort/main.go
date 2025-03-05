package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		chave := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > chave {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = chave
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Vetor original:", arr)
	insertionSort(arr)
	fmt.Println("Vetor ordenado:", arr)
}
