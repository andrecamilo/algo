package main

import "fmt"

/*
Algoritmo KMP (Knuth-Morris-Pratt)
O KMP é um algoritmo eficiente para busca de um padrão dentro de um texto.
Primeiro, calcula-se o array de LPS (Longest Prefix Suffix) e depois realiza a busca.
*/

// Calcula o array LPS para o padrão
func computeLPSArray(padrao string) []int {
	n := len(padrao)
	lps := make([]int, n)
	length := 0 // comprimento do prefixo anterior
	i := 1

	for i < n {
		if padrao[i] == padrao[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func KMPSearch(texto, padrao string) []int {
	lps := computeLPSArray(padrao)
	var resultados []int

	i, j := 0, 0
	n, m := len(texto), len(padrao)
	for i < n {
		if padrao[j] == texto[i] {
			i++
			j++
		}

		if j == m {
			resultados = append(resultados, i-j)
			j = lps[j-1]
		} else if i < n && padrao[j] != texto[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return resultados
}

func main() {
	texto := "ABABDABACDABABCABAB"
	padrao := "ABABCABAB"
	posicoes := KMPSearch(texto, padrao)
	fmt.Printf("Padrão encontrado nas posições: %v\n", posicoes)
}
