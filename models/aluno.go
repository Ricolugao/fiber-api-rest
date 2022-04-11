package models

import "time"

type Aluno struct {
	Id           int       `json:"id"`
	Nome         string    `json:"nome"`
	CPF          string    `json:"cpf"`
	RG           string    `json:"rg"`
	Data_criacao time.Time `json:"data_criacao"`
}

var Alunos []Aluno
