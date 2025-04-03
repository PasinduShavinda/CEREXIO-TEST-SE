
package models

type InventoryOverall struct {
	ProductId              int      `json:"product_id"`
	ProductName            *string  `json:"product_name"`
	Category               *string  `json:"category"`
	Brand                  *string  `json:"brand"`
	StockInHand            *int     `json:"stock_in_hand"`
	UnitPrice              *float64 `json:"unit_price"`
	SupplierId             *int     `json:"supplier_id"`
	DateAdded              *string  `json:"date_added"`
	LastUpdated            *string  `json:"last_updated"`
	Remarks                *string  `json:"remarks"`
	TransactionId          int      `json:"transaction_id"`
	TransactionType        *string  `json:"transaction_type"`
	Quantity               *int     `json:"quantity"`
	TransactionDate        *string  `json:"transaction_date"`
	TransactionDescription *string  `json:"transaction_description"`
	TransactionAmount      *float64 `json:"transaction_amount"`
	TransactionStatus      *string  `json:"transaction_status"`
	TransactionNotes       *string  `json:"transaction_notes"`
}