package main

import (
	"fmt"
)

func funcao[T MinhasRestricoes](arg T) T {
	return arg
}

type MinhasRestricoes interface {
	int | ~string | float64 | bool
}

type minhaString string

func main() {
	var ms minhaString = "Luan"
	fmt.Println(funcao[int](10))
	fmt.Println(funcao[string]("Luan"))
	fmt.Println(funcao[float64](10.5))
	fmt.Println(funcao[bool](true))
	fmt.Println(funcao(ms))
	//constraints.Ordered = 10

}
