package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ricolugao/fiber-api-rest/database"
	"github.com/ricolugao/fiber-api-rest/models"
)

func Saudacao(c *fiber.Ctx) error {
	return c.SendString("{\"Api Diz:\", \"E ai " + c.Params("nome") + ", tudo beleza?\"}")
}

func ExibeTodosAlunos(c *fiber.Ctx) error {
	db := database.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from alunos")
	if err != nil {
		log.Fatal(err)
	}

	aluno := models.Aluno{}
	alunos := []models.Aluno{}

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
		aluno.Data_criacao = data_criacao

		alunos = append(alunos, aluno)

	}

	return c.Status(http.StatusOK).JSON(alunos)
}

func CriaNovoAluno(c *fiber.Ctx) error {
	db := database.ConectaComBancoDeDados()
	defer db.Close()
	aluno := models.Aluno{}
	err := c.BodyParser(&aluno)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	insertNovoAluno, err := db.Query("insert into alunos (nome, cpf, rg) values ($1, $2, $3) returning id, data_criacao", aluno.Nome, aluno.CPF, aluno.RG)

	if err != nil {
		panic(err.Error())
	}

	for insertNovoAluno.Next() {
		err = insertNovoAluno.Scan(&aluno.Id, &aluno.Data_criacao)
		if err != nil {
			panic(err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON(aluno)
}

func EditaAluno(c *fiber.Ctx) error {
	db := database.ConectaComBancoDeDados()
	defer db.Close()
	return c.SendString("Aqui vai ser o script de edição de aluno}")
}

func ExibeAlunoPorId(c *fiber.Ctx) error {
	db := database.ConectaComBancoDeDados()
	defer db.Close()
	id := c.Params("id")
	// db.Query("select * from alunos where id = $1", id)
	return c.SendString("Aluno Exibindo apenas o aluno " + id)
}
