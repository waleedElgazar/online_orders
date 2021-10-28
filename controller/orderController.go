package controller

import (
	"github.com/gofiber/fiber/v2"
)

func InsertOrder(ctx *fiber.Ctx) error {

	//database.AddOrder(order)
	AddOrderItem(ctx)
	//fmt.Println(CommonUserId, "--------------------------")
	//return ctx.JSON(order)
	return nil
}
