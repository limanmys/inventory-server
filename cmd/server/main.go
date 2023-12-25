package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/routes"
	"github.com/limanmys/inventory-server/internal/migrations"
	"github.com/limanmys/inventory-server/internal/server"
)

func main() {
	if !fiber.IsChild() {
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	// Mount routes
	routes.Routes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
