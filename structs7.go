package main

import "fmt"

type Animal struct {
	Nome    string
	Espécie string
	Idade   int
	Som     string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a *Animal) ImprimirInformações() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Espécie)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Som: %s\n", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Rex",
		Espécie: "Cachorro",
		Idade:   5,
		Som:     "Au Au",
	}

	animal.ImprimirInformações()

	animal.ModificarSom("Woof Woof")

	animal.ImprimirInformações()
}
