package models

import (
	"log"
	"time"

	"github.com/ricolugao/fiber-api-rest/database"
)

type Aluno struct {
	Id          int       `json:"id"`
	Nome        string    `json:"nome"`
	CPF         string    `json:"cpf"`
	RG          string    `json:"rg"`
	DataCriacao time.Time `json:"data_criacao"`
}

func BuscaTodosAlunos() []Aluno {
	db := database.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from alunos")
	if err != nil {
		log.Fatal(err)
	}

	aluno := Aluno{}
	alunos := []Aluno{}

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome, cpf, rg string
		var data_criacao time.Time

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &cpf, &rg, &data_criacao)
		if err != nil {
			log.Fatal(err)
		}

		aluno.Id = id
		aluno.Nome = nome
		aluno.CPF = cpf
		aluno.RG = rg
		aluno.DataCriacao = data_criacao

		alunos = append(alunos, aluno)

	}

	return alunos
}

func InsereNovoAluno(aluno Aluno) Aluno {
	var AlunoIserido Aluno
	AlunoIserido.Nome = aluno.Nome
	AlunoIserido.CPF = aluno.CPF
	AlunoIserido.RG = aluno.RG

	db := database.ConectaComBancoDeDados()
	defer db.Close()
	insertNovoAluno, err := db.Query("insert into alunos (nome, cpf, rg) values ($1, $2, $3) returning id, data_criacao", aluno.Nome, aluno.CPF, aluno.RG)
	if err != nil {
		panic(err.Error())
	}

	for insertNovoAluno.Next() {
		err = insertNovoAluno.Scan(&AlunoIserido.Id, &AlunoIserido.DataCriacao)
		if err != nil {
			panic(err.Error())
		}
	}
	return AlunoIserido
}

var Alunos []Aluno
