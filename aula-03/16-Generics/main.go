package main

import "fmt"

// Estou comparando o valor do slice com o item e não slice contra slice
func Contains[T comparable](slice []T, item T) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}

	return false
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	nomes := []string{"Luan", "Ana", "Carlos"}
	booleanos := []bool{true, false}

	fmt.Println(Contains(numeros, 3))      // true
	fmt.Println(Contains(numeros, 10))     // false
	fmt.Println(Contains(nomes, "Luan"))   // true
	fmt.Println(Contains(nomes, "Maria"))  // false
	fmt.Println(Contains(booleanos, true)) // true
}
