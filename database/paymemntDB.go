package database

import (
	"fmt"

	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/models"
)

func AddPaymentForOrder(payment models.Payment) error{
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "INSERT INTO newResturant.payment set userId=?, orderId=?, paymentDate=?, amount=?,paymentType=?"
	insert, err := DB.Prepare(query)
	if err!=nil {
		fmt.Println("error while inserting payment data",err)
		return err
	}
	_,err=insert.Exec(payment.UserId,payment.OrderId,payment.PaymentDate,payment.Amount,payment.PaymentType)
	if err!=nil {
		fmt.Println("error while inserting the data",err)
		return err
	}
	return nil
}	
