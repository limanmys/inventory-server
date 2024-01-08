package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/limanmys/inventory-server/pkg/aes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/limanmys/inventory-server/app/routes"
	"github.com/limanmys/inventory-server/internal/migrations"
	"github.com/limanmys/inventory-server/internal/seeds"
	"github.com/limanmys/inventory-server/internal/server"
)

func main() {
	// Migrate tables
	if !fiber.IsChild() {
		//Migrate tables
		if err := migrations.Migrate(); err != nil {
			log.Println("error when migrating tables, ", err.Error())
		}

		// Seed alternative packages
		seeds.Init()
	}

	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	// Add logger
	app.Use(logger.New())

	// Add compress
	app.Use(compress.New())

	// Add recover with stack tracing
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	// Mount routes
	routes.Routes(app)

	// Start server
	log.Fatal(app.Listen("127.0.0.1:7806"))
}
