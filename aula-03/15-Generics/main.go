package main

import "fmt"

func funcao[T comparable](arg T) T {
	return arg
}

func main() {
	fmt.Println(funcao[int](10))
	fmt.Println(funcao[string]("Luan"))
	fmt.Println(funcao[float64](10.5))
	fmt.Println(funcao[bool](true))
	fmt.Println(funcao[[]int]([]int{1, 2, 3}))
}
