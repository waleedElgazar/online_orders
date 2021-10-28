package controller

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/database"
	"github.com/waleedElgazar/resturant/models"
	"gopkg.in/mail.v2"
)

func PayForOrder(ctx *fiber.Ctx) error {
	var data models.Payment
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error while parsing data", err)
		return err
	}
	currentTime := time.Now()
	payment := models.Payment{
		PaymentDate: currentTime.Format("2006.01.02 15:04:05"),
		PaymentType: data.PaymentType,
		UserId:      data.UserId,
		OrderId:     data.OrderId,
		Amount:      data.Amount,
	}
	err = database.AddPaymentForOrder(payment)
	SendPaymentNotification("walidreda428@gmail.com",payment)
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return ctx.JSON(
			fiber.Map{
				"message": "order id or user id are not found",
			},
		)
	}
	return ctx.JSON(data)
}

func SendPaymentNotification(UserAccount string,data models.Payment) {
	adc := mail.NewMessage()
	adc.SetHeader("From", os.Getenv("EMAIL"))
	adc.SetHeader("To", UserAccount)
	adc.SetHeader("Subject", "hi from golang")
	amount:=fmt.Sprintf("%f", data.Amount)
	x:="we noticed that you pay "+amount+"for the order with id "+strconv.Itoa(int(data.OrderId))
	adc.SetBody("text/plain", x)
	a := mail.NewDialer("smtp.gmail.com", 587, "walidreda427@gmail.com", os.Getenv("PASSWORD"))
	if err := a.DialAndSend(adc); err != nil {
		fmt.Println("error ", err)
		panic(err)
	}
}
