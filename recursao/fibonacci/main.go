package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivo := arr[len(arr)/2]
	var menores, iguais, maiores []int

	for _, num := range arr {
		if num < pivo {
			menores = append(menores, num)
		} else if num == pivo {
			iguais = append(iguais, num)
		} else {
			maiores = append(maiores, num)
		}
	}
	menores = quickSort(menores)
	maiores = quickSort(maiores)
	return append(append(menores, iguais...), maiores...)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Vetor original:", arr)
	ordenado := quickSort(arr)
	fmt.Println("Vetor ordenado:", ordenado)
}
