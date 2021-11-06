package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/routes"
	"log"
)
func welcome(ctx *fiber.Ctx) error {
	ctx.Write([]byte("hello"))
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configration.OpenConnection()
	app := fiber.New()
	app.Post("/",welcome)
	routes.UserSetUp(app)
	routes.OrderSetUp(app)
	routes.PaymentSetUp(app)
	app.Listen(":8000")
}