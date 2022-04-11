package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ricolugao/fiber-api-rest/controllers"
)

func HandleRequests() {
	app := Setup()
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Get("/alunos", controllers.ExibeTodosAlunos)
	app.Post("/alunos", controllers.Insert)
	app.Get("/alunos/:id", controllers.ExibeAlunoPorId)
	app.Patch("/alunos", controllers.EditaAluno)
	app.Get("/:nome", controllers.Saudacao)

	log.Fatal(app.Listen(":8080"))
}

func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Return the configured app
	return app
}
