package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/PasinduShavinda/go-project-api/internal/db"
	"github.com/PasinduShavinda/go-project-api/internal/models"
)

// Define API methods - Inventory Overall

// OK Tested
// @Summary Fetch overall inventory details
// @Description Get a comprehensive list of inventory and transaction details
// @Tags inventory_overall
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} models.InventoryOverall
// @Failure 500 {object} map[string]string
// @Router /inventory-details-overall [get]
func GetOverallInventoryDetails(w http.ResponseWriter, r *http.Request) {
	var inv_ov []models.InventoryOverall
	rows, err := db.DB.Query(
		`SELECT 
        t1.product_id, 
        t1.product_name, 
        t1.category, 
        t1.brand, 
        t1.stock_in_hand,
        t1.unit_price, 
        t1.supplier_id, 
        t1.date_added, 
        t1.last_updated, 
        t1.remarks,
        t2.transaction_id, 
        t2.transaction_type, 
        t2.quantity, 
        t2.transaction_date, 
        t2.transaction_description, 
        t2.transaction_amount, 
        t2.transaction_status,
        t2.transaction_notes 
        FROM inventory_master t1 
        INNER JOIN inventory_details t2
        ON t1.product_id = t2.product_id`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var invov models.InventoryOverall
		if err := rows.Scan(
			&invov.ProductId,
			&invov.ProductName,
			&invov.Category,
			&invov.Brand,
			&invov.StockInHand,
			&invov.UnitPrice,
			&invov.SupplierId,
			&invov.DateAdded,
			&invov.LastUpdated,
			&invov.Remarks,
			&invov.TransactionId,
			&invov.TransactionType,
			&invov.Quantity,
			&invov.TransactionDate,
			&invov.TransactionDescription,
			&invov.TransactionAmount,
			&invov.TransactionStatus,
			&invov.TransactionNotes,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_ov = append(inv_ov, invov)
	}
	json.NewEncoder(w).Encode(inv_ov)
}

// @Summary Filter overall inventory details
// @Description Retrieve overall inventory records based on query parameters
// @Tags inventory_overall
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param product_id query string false "Product ID"
// @Param product_name query string false "Product Name"
// @Param category query string false "Category"
// @Param brand query string false "Brand"
// @Param stock_in_hand query string false "Stock In Hand"
// @Param unit_price query string false "Unit Price"
// @Param supplier_id query string false "Supplier ID"
// @Param date_added query string false "Date Added"
// @Param last_updated query string false "Last Updated"
// @Param remarks query string false "Remarks"
// @Param transaction_id query string false "Transaction ID"
// @Param transaction_type query string false "Transaction Type"
// @Param quantity query string false "Quantity"
// @Param transaction_date query string false "Transaction Date"
// @Param transaction_description query string false "Transaction Description"
// @Param transaction_amount query string false "Transaction Amount"
// @Param transaction_status query string false "Transaction Status"
// @Param transaction_notes query string false "Transaction Notes"
// @Success 200 {array} models.InventoryOverall
// @Failure 500 {object} map[string]string
// @Router /inventory-details-overall-filter [get]
func GetFilteredOverallInventoryDetails(w http.ResponseWriter, r *http.Request) {

	query :=
		`SELECT 
        t1.product_id, 
        t1.product_name, 
        t1.category, 
        t1.brand, 
        t1.stock_in_hand,
        t1.unit_price, 
        t1.supplier_id, 
        t1.date_added, 
        t1.last_updated, 
        t1.remarks,
        t2.transaction_id, 
        t2.transaction_type, 
        t2.quantity, 
        t2.transaction_date, 
        t2.transaction_description, 
        t2.transaction_amount, 
        t2.transaction_status,
        t2.transaction_notes 
        FROM inventory_master t1 
        INNER JOIN inventory_details t2
        ON t1.product_id = t2.product_id WHERE 1=1`
	args := []interface{}{}

	if product_id := r.URL.Query().Get("product_id"); product_id != "" {
		query += " AND t1.product_id = ?"
		args = append(args, product_id)
	}

	if product_name := r.URL.Query().Get("product_name"); product_name != "" {
		query += " AND t1.product_name = ?"
		args = append(args, product_name)
	}

	if category := r.URL.Query().Get("category"); category != "" {
		query += " AND t1.category = ?"
		args = append(args, category)
	}

	if brand := r.URL.Query().Get("brand"); brand != "" {
		query += " AND t1.brand = ?"
		args = append(args, brand)
	}

	if stock_in_hand := r.URL.Query().Get("stock_in_hand"); stock_in_hand != "" {
		query += " AND t1.stock_in_hand = ?"
		args = append(args, stock_in_hand)
	}

	if unit_price := r.URL.Query().Get("unit_price"); unit_price != "" {
		query += " AND t1.unit_price = ?"
		args = append(args, unit_price)
	}

	if supplier_id := r.URL.Query().Get("supplier_id"); supplier_id != "" {
		query += " AND t1.supplier_id = ?"
		args = append(args, supplier_id)
	}

	if date_added := r.URL.Query().Get("date_added"); date_added != "" {
		query += " AND t1.date_added = ?"
		args = append(args, date_added)
	}

	if last_updated := r.URL.Query().Get("last_updated"); last_updated != "" {
		query += " AND t1.last_updated = ?"
		args = append(args, last_updated)
	}

	if remarks := r.URL.Query().Get("remarks"); remarks != "" {
		query += " AND t1.remarks = ?"
		args = append(args, remarks)
	}

	if transaction_id := r.URL.Query().Get("transaction_id"); transaction_id != "" {
		query += " AND t2.transaction_id = ?"
		args = append(args, transaction_id)
	}

	if transaction_type := r.URL.Query().Get("transaction_type"); transaction_type != "" {
		query += " AND t2.transaction_type = ?"
		args = append(args, transaction_type)
	}

	if quantity := r.URL.Query().Get("quantity"); quantity != "" {
		query += " AND t2.quantity = ?"
		args = append(args, quantity)
	}

	if transaction_date := r.URL.Query().Get("transaction_date"); transaction_date != "" {
		query += " AND t2.transaction_date = ?"
		args = append(args, transaction_date)
	}

	if transaction_description := r.URL.Query().Get("transaction_description"); transaction_description != "" {
		query += " AND t2.transaction_description = ?"
		args = append(args, transaction_description)
	}

	if transaction_amount := r.URL.Query().Get("transaction_amount"); transaction_amount != "" {
		query += " AND t2.transaction_amount = ?"
		args = append(args, transaction_amount)
	}

	if transaction_status := r.URL.Query().Get("transaction_status"); transaction_status != "" {
		query += " AND t2.transaction_status = ?"
		args = append(args, transaction_status)
	}

	if transaction_notes := r.URL.Query().Get("transaction_notes"); transaction_notes != "" {
		query += " AND t2.transaction_notes = ?"
		args = append(args, transaction_notes)
	}

	rows, err := db.DB.Query(query, args...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inv_o []models.InventoryOverall
	for rows.Next() {
		var invo models.InventoryOverall
		err := rows.Scan(
			&invo.ProductId,
			&invo.ProductName,
			&invo.Category,
			&invo.Brand,
			&invo.StockInHand,
			&invo.UnitPrice,
			&invo.SupplierId,
			&invo.DateAdded,
			&invo.LastUpdated,
			&invo.Remarks,
			&invo.TransactionId,
			&invo.TransactionType,
			&invo.Quantity,
			&invo.TransactionDate,
			&invo.TransactionDescription,
			&invo.TransactionAmount,
			&invo.TransactionStatus,
			&invo.TransactionNotes,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_o = append(inv_o, invo)
	}

	json.NewEncoder(w).Encode(inv_o)
}