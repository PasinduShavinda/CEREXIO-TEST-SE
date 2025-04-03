package models

// Inventory Details Model
type InventoryDetails struct {
	TransactionId          int      `json:"transaction_id"`
	ProductId              *int     `json:"product_id"`
	TransactionType        *string  `json:"transaction_type"`
	Quantity               *int     `json:"quantity"`
	TransactionDate        *string  `json:"transaction_date"`
	TransactionDescription *string  `json:"transaction_description"`
	TransactionAmount      *float64 `json:"transaction_amount"`
	TransactionStatus      *string  `json:"transaction_status"`
	TransactionNotes       *string  `json:"transaction_notes"`
}