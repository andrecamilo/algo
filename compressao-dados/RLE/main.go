package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rleEncode(input string) string {
	if len(input) == 0 {
		return ""
	}
	var encoded strings.Builder
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			encoded.WriteByte(input[i-1])
			encoded.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	// Processa o último grupo
	encoded.WriteByte(input[len(input)-1])
	encoded.WriteString(strconv.Itoa(count))
	return encoded.String()
}

func rleDecode(input string) string {
	var decoded strings.Builder
	n := len(input)
	for i := 0; i < n; i++ {
		ch := input[i]
		j := i + 1
		countStr := ""
		for j < n && input[j] >= '0' && input[j] <= '9' {
			countStr += string(input[j])
			j++
		}
		var count int
		fmt.Sscanf(countStr, "%d", &count)
		for k := 0; k < count; k++ {
			decoded.WriteByte(ch)
		}
		i = j - 1
	}
	return decoded.String()
}

func main() {
	original := "aaaabbbccaa"
	encoded := rleEncode(original)
	decoded := rleDecode(encoded)
	fmt.Println("Original:", original)
	fmt.Println("Codificado:", encoded)
	fmt.Println("Decodificado:", decoded)
}
