package main

import (
	"fmt"
	"strings"
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
	playlist1 := Playlist{
		Nome: "Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: time.Minute*3 + time.Second*45},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: time.Minute*4 + time.Second*12},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: time.Minute*2 + time.Second*58},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: time.Minute*3 + time.Second*15},
			{Titulo: "Música 5", Artista: "Artista 5", Duracao: time.Minute*2 + time.Second*30},
			{Titulo: "Música 6", Artista: "Artista 6", Duracao: time.Minute*4 + time.Second*5},
		},
	}

	playlist3 := Playlist{
		Nome: "Playlist 3",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 7", Duracao: time.Minute*3 + time.Second*45},
			{Titulo: "Música 7", Artista: "Artista 8", Duracao: time.Minute*2 + time.Second*20},
			{Titulo: "Música 8", Artista: "Artista 9", Duracao: time.Minute*3 + time.Second*50},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	titulo := "Música 1"
	resultado := buscarPlaylistsPorTitulo(titulo, playlists)

	if len(resultado) == 0 {
		fmt.Println("Nenhuma playlist encontrada com a música:", titulo)
	} else {
		fmt.Println("Playlists com a música", titulo+":")
		for _, playlist := range resultado {
			fmt.Println("Nome:", playlist.Nome)
		}
	}
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if strings.EqualFold(musica.Titulo, titulo) {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}
