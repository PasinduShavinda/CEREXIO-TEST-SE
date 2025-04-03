package models

// Inventory Master Model
type InventoryMaster struct {
	ProductId   int      `json:"product_id"`
	ProductName *string  `json:"product_name"`
	Category    *string  `json:"category"`
	Brand       *string  `json:"brand"`
	StockInHand *int     `json:"stock_in_hand"`
	UnitPrice   *float64 `json:"unit_price"`
	SupplierId  *int     `json:"supplier_id"`
	DateAdded   *string  `json:"date_added"`
	LastUpdated *string  `json:"last_updated"`
	Remarks     *string  `json:"remarks"`
}
