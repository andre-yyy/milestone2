package dto

type CreateTopupDto struct {
	ExternalID string `json:"external_id"`
	InvoiceID  string `json:"invoice_id"`
}
