package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}
	return float64(soma) / float64(len(f.Avaliacoes))
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Println("Título:", f.Titulo)
	fmt.Println("Diretor:", f.Diretor)
	fmt.Println("Ano:", f.Ano)
	fmt.Println("Média de Avaliações:", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "Inception",
		Diretor:    "Christopher Nolan",
		Ano:        2010,
		Avaliacoes: []int{8, 9, 10},
	}

	filme.ImprimirInformacoes()

	filme.AdicionarAvaliacao(7)
	filme.RemoverAvaliacao(1)

	filme.ImprimirInformacoes()
}
