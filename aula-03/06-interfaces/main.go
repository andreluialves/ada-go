package main

import "fmt"

type Animal interface {
	FazerSom() string
	Andar() string
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) FazerSom() string {
	return "au au"
}

func (c Cachorro) Andar() string {
	return c.Nome + " esta andando pelo quintal"
}

type Gato struct {
	Nome string
}

func (g Gato) FazerSom() string {
	return "miau"
}

func (g Gato) Andar() string {
	return g.Nome + " esta andando em cima do muro"
}

func apresentarAnimal(animal Animal) {
	fmt.Println("som:", animal.FazerSom())
	fmt.Println("movimento:", animal.Andar())
}

func main() {
	cachorro := Cachorro{Nome: "Rex"}
	gato := Gato{Nome: "Mimi"}

	apresentarAnimal(cachorro)
	apresentarAnimal(gato)
}
