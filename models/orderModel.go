package models

type Order struct {
	OrderID 	uint		`json:"id"`
	UserId		uint		`json:"user_id"`
	TotalPrice	float64		`json:"total_price"`
}		