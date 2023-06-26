package main

import (
	"fmt"
	"time"
)

type Funcionário struct {
	Nome    string
	Salário float64
	Idade   int
}

func (f *Funcionário) AumentarSalário(porcentagem float64) {
	aumento := f.Salário * (porcentagem / 100)
	f.Salário += aumento
}

func (f *Funcionário) DiminuirSalário(porcentagem float64) {
	diminuição := f.Salário * (porcentagem / 100)
	f.Salário -= diminuição
}

func (f *Funcionário) TempoDeServiço() int {
	idadeInicio := 18
	anosDeServiço := time.Now().Year() - (idadeInicio + f.Idade)
	return anosDeServiço
}

func main() {
	funcionário := Funcionário{
		Nome:    "João",
		Salário: 5000.0,
		Idade:   30,
	}

	fmt.Println("Salário inicial:", funcionário.Salário)
	funcionário.AumentarSalário(10)
	fmt.Println("Salário após aumento de 10%:", funcionário.Salário)
	funcionário.DiminuirSalário(5)
	fmt.Println("Salário após diminuição de 5%:", funcionário.Salário)

	tempoDeServiço := funcionário.TempoDeServiço()
	fmt.Println("Tempo de serviço:", tempoDeServiço, "anos")
}
