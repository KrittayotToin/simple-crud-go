package main

import (
	"fmt"
	"os"

	"github.com/KrittayotToin/simple-crud-go/internal/config"
	"github.com/KrittayotToin/simple-crud-go/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := config.ConnectMongoDB(); err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}
	defer config.DisconnectMongoDB()

	// Routes
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	routes.SetRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on 😂 :%s\n", port)
	app.Listen(":" + port)
}
