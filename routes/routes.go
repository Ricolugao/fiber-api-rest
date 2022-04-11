package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ricolugao/fiber-api-rest/controllers"
)

func HandleRequests() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Get("/alunos", controllers.ExibeTodosAlunos)
	app.Post("/alunos", controllers.CriaNovoAluno)
	app.Get("/alunos/:id", controllers.ExibeAlunoPorId)
	app.Patch("/alunos", controllers.EditaAluno)
	app.Get("/:nome", controllers.Saudacao)

	log.Fatal(app.Listen(":8080"))
}
