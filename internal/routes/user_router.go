package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Post("/users", createUserHandler)
	app.Get("/users/:id", getUserByIDHandler)
}

func createUserHandler(c *fiber.Ctx) error {
	fmt.Println("createUserHandler function called")
	return c.SendString("Create User")
}

func getUserByIDHandler(c *fiber.Ctx) error {
	userID := c.Params("id")
	fmt.Printf("getUserByIDHandler function called with ID: %s\n", userID)
	return c.SendString("Get User by ID: " + userID)
}
