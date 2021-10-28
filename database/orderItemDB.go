package database

import (
	"fmt"

	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/models"
)

func AddOderItem(orderItem models.OrderItem) {
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "INSERT INTO newResturant.orderItem set orderId=?, itemName=?, itemQuantity=?, itemPrice=?"
	insert, err := DB.Prepare(query)
	if err != nil {
		fmt.Println("error while executing insert query",err)
		return 
	}
	_, err = insert.Exec(orderItem.OrderID,orderItem.ItemName,orderItem.ItemQuantity,orderItem.ItemPrice)
	if err != nil {
		fmt.Println("error while parsing inserting data itemorder",err)
		
	}

}
