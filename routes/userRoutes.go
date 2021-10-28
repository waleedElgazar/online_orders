package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/controller"
)

func UserSetUp(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Post("/verify", controller.VerifyAccount)
	app.Get("/user",controller.GetUser)
	app.Get("/logout",controller.LogOut)
}
