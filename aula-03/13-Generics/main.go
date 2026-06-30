package main

import "fmt"

// Generics, vira um tipo concreto
// Em tempo de compilação
func funcao[T any](arg T) T {
	return arg
}

// Interface vazia, o tipo é definido em tempo de execução(runtime)
// Envolve type assertions e type switches
func funcao2(arg any) any {
	return arg
}

func main() {
	fmt.Println(funcao[int](10))
	fmt.Println(funcao[string]("Luan"))
	fmt.Println(funcao[float64](10.5))
	fmt.Println(funcao[bool](true))
}
