package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limanmys/inventory-server/app/controllers/discoveries"
)

func Routes(app *fiber.App) {
	app.Get("/", discoveries.Create)
}
