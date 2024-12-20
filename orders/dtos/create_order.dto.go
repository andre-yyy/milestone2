package dto

type CreateOrderDto struct {
	ProductID int `json:"product_id"`
	RentDays  int `json:"rent_days"`
}
