package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/database"
	"github.com/waleedElgazar/resturant/models"
)

func AddOrderItem(ctx *fiber.Ctx) error {
	if IsAuthorized(ctx) != nil {
		return ctx.JSON(fiber.Map{
			"message": "unable to order as u didn't login",
		})
	} else {
		var data []models.OrderItem
		err := ctx.BodyParser(&data)
		if err != nil {
			fmt.Println("error while parsing data", err)
			return err
		}
		var orderItems []models.OrderItem
		order := models.Order{
			UserId:     uint(CommonUserId),
			TotalPrice: 10,
		}
		IdOrder,_:=database.AddOrder(order)
		n:=len(data)
		var totalPrice float64
		for i := 0; i <n; i++ {
			order:=models.OrderItem{
				OrderID: uint(IdOrder),
				ItemName: data[i].ItemName,
				ItemQuantity: data[i].ItemQuantity,
				ItemPrice: data[i].ItemPrice,
			}
			totalPrice+=order.ItemPrice*float64(order.ItemQuantity)
			database.AddOderItem(order)
			orderItems=append(orderItems,order)
		}
		database.UpdatePrice(int(IdOrder),totalPrice)
		return ctx.JSON(orderItems)
	}
}
