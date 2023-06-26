package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func main() {
	musicas := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: time.Minute*3 + time.Second*45},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: time.Minute*4 + time.Second*12},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: time.Minute*2 + time.Second*58},
	}

	playlist := Playlist{
		Nome:    "Minha Playlist",
		Musicas: musicas,
	}

	imprimirPlaylist(playlist)
}

func imprimirPlaylist(p Playlist) {
	fmt.Println("Nome da Playlist:", p.Nome)

	var duracaoTotal time.Duration

	for _, musica := range p.Musicas {
		fmt.Println("Título:", musica.Titulo)
		fmt.Println("Artista:", musica.Artista)
		fmt.Println("Duração:", musica.Duracao)
		fmt.Println()

		duracaoTotal += musica.Duracao
	}

	fmt.Println("Tempo total da Playlist:", duracaoTotal)
}
