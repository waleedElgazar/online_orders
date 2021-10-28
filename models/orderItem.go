package models

type OrderItem struct {
	Id				uint  	`json:"id"`
	OrderID			uint	`json:"order_id"`
	ItemName		string	`json:"item_name"`
	ItemQuantity	int		`json:"item_quantity"`
	ItemPrice		float64	`json:"item_price"`
}