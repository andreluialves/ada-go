package main

import "fmt"

type Animal interface {
	FazerSom() string
}

type Cachorro struct {
	Nome string
}

func (c *Cachorro) FazerSom() string {
	return "au au"
}

func (c *Cachorro) BuscarBolinha() string {
	return c.Nome + " buscou a bolinha"
}

type Gato struct {
	Nome string
}

func (g *Gato) FazerSom() string {
	return "miau"
}

func (g *Gato) SubirNoMuro() string {
	return g.Nome + " subiu no muro"
}

func fazerTypeAssertion(animal Animal) {
	cachorro, ok := animal.(*Cachorro) // Type assertion
	if ok {
		fmt.Println("type assertion: recebi um cachorro")
		fmt.Println(cachorro.BuscarBolinha())
		return
	}

	fmt.Println("type assertion: nao recebi um cachorro")
}

func fazerTypeSwitch(animal Animal) {
	switch animalConvertido := animal.(type) {
	case *Cachorro:
		fmt.Println("type switch: cachorro")
		fmt.Println("som:", animalConvertido.FazerSom())
		fmt.Println(animalConvertido.BuscarBolinha())
	case *Gato:
		fmt.Println("type switch: gato")
		fmt.Println("som:", animalConvertido.FazerSom())
		fmt.Println(animalConvertido.SubirNoMuro())
	default:
		fmt.Println("type switch: animal desconhecido")
	}
}

func main() {
	cachorro := &Cachorro{Nome: "Rex"}
	gato := &Gato{Nome: "Mimi"}

	fazerTypeAssertion(cachorro)
	fazerTypeAssertion(gato)

	fazerTypeSwitch(cachorro)
	fazerTypeSwitch(gato)
}
