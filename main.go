package main

import (
	"github.com/Alencars/api-go-gin/models"
	"github.com/Alencars/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "João Vitor", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "4800000000"},
	}
	routes.HandleRequests()
}
