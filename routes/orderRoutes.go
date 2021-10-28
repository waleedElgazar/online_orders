package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/controller"
)

func OrderSetUp(app *fiber.App) {
	app.Post("/addorder", controller.InsertOrder)
}
