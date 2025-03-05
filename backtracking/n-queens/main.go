package main

import "fmt"

// Problema das N Rainhas (N-Queens)
// Resolvemos a colocação de N rainhas em um tabuleiro N×N de modo que nenhuma rainha se ataque.
// Cada posição segura é verificada recursivamente.

// Verifica se é seguro colocar uma rainha na coluna 'col' da linha 'row'
func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		// Verifica mesmo coluna e diagonais
		if board[i] == col || board[i]-i == col-row || board[i]+i == col+row {
			return false
		}
	}
	return true
}

func solveNQueensUtil(n, row int, board []int, solutions *[][]int) {
	if row == n {
		solution := make([]int, n)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}
	for col := 0; col < n; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveNQueensUtil(n, row+1, board, solutions)
		}
	}
}

func solveNQueens(n int) [][]int {
	var solutions [][]int
	board := make([]int, n)
	solveNQueensUtil(n, 0, board, &solutions)
	return solutions
}

func main() {
	n := 8
	solutions := solveNQueens(n)
	fmt.Printf("Número de soluções para %d rainhas: %d\n", n, len(solutions))
	// Imprime a primeira solução (cada índice representa a coluna da rainha em cada linha)
	if len(solutions) > 0 {
		fmt.Println("Primeira solução:", solutions[0])
	}
}
