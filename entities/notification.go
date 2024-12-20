package entity

import "time"

type Notification struct {
	ID                     string    `json:"id"`
	AdjustedReceivedAmount float64   `json:"adjusted_received_amount"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	BankCode               string    `json:"bank_code"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 float64   `json:"amount"`
	PaidAt                 time.Time `json:"paid_at"`
	PaidAmount             float64   `json:"paid_amount"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	Created                string    `json:"created"`
	Updated                string    `json:"updated"`
	Currency               string    `json:"currency"`
	PaymentDestination     string    `json:"payment_destination"`
}
