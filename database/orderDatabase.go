package database

import (
	"fmt"

	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/models"
)

func AddOrder(order models.Order)(int64,error) {
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "INSERT INTO newResturant.orders set userId=?"
	insert, err := DB.Prepare(query)
	if err != nil {
		fmt.Println("error while executing insert query")
		return -1,err
	}
	res, err := insert.Exec(order.UserId)
	if err != nil {
		fmt.Println("error while parsing inserting data order",err)
		return -1,err
	}
	return res.LastInsertId()
}

func UpdatePrice(order_id int,total_price float64){
	DB = configration.OpenConnection()
	defer DB.Close()
	query:="UPDATE newResturant.orders set totalPrice=? WHERE orderId=?"
	insert, err := DB.Prepare(query)
	if err != nil {
		fmt.Println("error while executing insert query")
		
	}
	_, err = insert.Exec(total_price,order_id)
	if err != nil {
		fmt.Println("error while parsing inserting data order",err)
		
	}
}
