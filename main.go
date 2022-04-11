package main

import (
	"github.com/gofiber/fiber"
	"github.com/ricolugao/fiber-api-rest/routes"
)

//{"id":"1","nome":"Lugão","CPF":"12345678910","RG":"12345678"}
func main() {
	routes.HandleRequests()
}

func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200

	// Return the configured app
	return app
}
