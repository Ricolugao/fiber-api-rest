package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricolugao/fiber-api-rest/database"
	"github.com/ricolugao/fiber-api-rest/models"
)

func Saudacao(c *fiber.Ctx) error {
	return c.SendString("{\"Api Diz:\", \"E ai " + c.Params("nome") + ", tudo beleza?\"}")
}

func ExibeTodosAlunos(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.BuscaTodosAlunos())
}

func Insert(c *fiber.Ctx) error {
	aluno := models.Aluno{}
	err := c.BodyParser(&aluno)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	insertNovoAluno := models.InsereNovoAluno(aluno)

	return c.Status(http.StatusOK).JSON(insertNovoAluno)
}

func ExcluiAluno(c *fiber.Ctx) error {
	db := database.ConectaComBancoDeDados()
	defer db.Close()
	// id := c.Params("id")

	// excluirAluno, err := db.Query("delete from alunos where id = $1", id)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.SendString("Aqui vai ser o script de excluir aluno}")
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
