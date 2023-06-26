package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}
	return soma / float64(len(a.Notas))
}

func (a *Aluno) ImprimirInformacoes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Média:", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{7.5, 8.0, 6.5},
	}

	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(7.0)
	aluno.RemoverNota(1)

	aluno.ImprimirInformacoes()
}
