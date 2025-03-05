package main

import "fmt"

var memo = make(map[int]int)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, existe := memo[n]; existe {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}
