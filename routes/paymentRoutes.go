package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/controller"
)

func PaymentSetUp(app *fiber.App) {
	app.Post("/addpayment",controller.PayForOrder)
}
