package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/PasinduShavinda/go-project-api/internal/db"
	"github.com/PasinduShavinda/go-project-api/internal/models"
	"github.com/gorilla/mux"
)

// Define API methods - Inventory Details

// OK Tested
// @Summary Fetch all inventory details records
// @Description Get a list of all inventory details records
// @Tags inventory_details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} models.InventoryDetails
// @Failure 500 {object} map[string]string
// @Router /inventory-details [get]
func GetAllInventoryDetails(w http.ResponseWriter, r *http.Request) {
	var inv_d []models.InventoryDetails
	rows, err := db.DB.Query("SELECT * FROM inventory_details")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var invd models.InventoryDetails
		if err := rows.Scan(
			&invd.TransactionId,
			&invd.ProductId,
			&invd.TransactionType,
			&invd.Quantity,
			&invd.TransactionDate,
			&invd.TransactionDescription,
			&invd.TransactionAmount,
			&invd.TransactionStatus,
			&invd.TransactionNotes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_d = append(inv_d, invd)
	}
	json.NewEncoder(w).Encode(inv_d)
}

// OK Tested
// @Summary Fetch inventory details by ID
// @Description Get a single inventory details record using its ID
// @Tags inventory_details\
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.InventoryDetails
// @Failure 404 {object} map[string]string
// @Router /inventory-details/{id} [get]
func GetInventoryDetailsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var invd models.InventoryDetails
	err := db.DB.QueryRow("SELECT * FROM inventory_details WHERE transaction_id = ?", id).Scan(
		&invd.ProductId,
		&invd.TransactionType,
		&invd.Quantity,
		&invd.TransactionDate,
		&invd.TransactionDescription,
		&invd.TransactionAmount,
		&invd.TransactionStatus,
		&invd.TransactionNotes)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(invd)
}

// OK Tested
// @Summary Filter inventory details records
// @Description Retrieve inventory details records based on query parameters
// @Tags inventory_details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param transaction_id query string false "Transaction ID"
// @Param product_id query string false "Product ID"
// @Param transaction_type query string false "Transaction Type"
// @Param quantity query string false "Quantity"
// @Param transaction_date query string false "Transaction Date"
// @Param transaction_description query string false "Transaction Description"
// @Param transaction_amount query string false "Transaction Amount"
// @Param transaction_status query string false "Transaction Status"
// @Param transaction_notes query string false "Transaction Notes"
// @Success 200 {array} models.InventoryDetails
// @Failure 500 {object} map[string]string
// @Router /inventory-details-filter [get]
func GetFilteredInventoryDetails(w http.ResponseWriter, r *http.Request) {
	
	query := "SELECT * FROM inventory_details WHERE 1=1"
	args := []interface{}{}

	if transaction_id := r.URL.Query().Get("transaction_id"); transaction_id != "" {
		query += " AND transaction_id = ?"
		args = append(args, transaction_id)
	}

	if product_id := r.URL.Query().Get("product_id"); product_id != "" {
		query += " AND product_id = ?"
		args = append(args, product_id)
	}

	if transaction_type := r.URL.Query().Get("transaction_type"); transaction_type != "" {
		query += " AND transaction_type = ?"
		args = append(args, transaction_type)
	}

	if quantity := r.URL.Query().Get("quantity"); quantity != "" {
		query += " AND quantity = ?"
		args = append(args, quantity)
	}

	if transaction_date := r.URL.Query().Get("transaction_date"); transaction_date != "" {
		query += " AND transaction_date = ?"
		args = append(args, transaction_date)
	}

	if transaction_description := r.URL.Query().Get("unit_price"); transaction_description != "" {
		query += " AND transaction_description = ?"
		args = append(args, transaction_description)
	}

	if transaction_amount := r.URL.Query().Get("transaction_amount"); transaction_amount != "" {
		query += " AND transaction_amount = ?"
		args = append(args, transaction_amount)
	}

	if transaction_status := r.URL.Query().Get("transaction_status"); transaction_status != "" {
		query += " AND transaction_status = ?"
		args = append(args, transaction_status)
	}

	if transaction_notes := r.URL.Query().Get("transaction_notes"); transaction_notes != "" {
		query += " AND transaction_notes = ?"
		args = append(args, transaction_notes)
	}

	rows, err := db.DB.Query(query, args...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inv_d []models.InventoryDetails
	for rows.Next() {
		var invd models.InventoryDetails
		err := rows.Scan(&invd.TransactionId, &invd.ProductId, &invd.TransactionType, &invd.Quantity,
			&invd.TransactionDate, &invd.TransactionDescription, &invd.TransactionAmount,
			&invd.TransactionStatus, &invd.TransactionNotes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_d = append(inv_d, invd)
	}

	json.NewEncoder(w).Encode(inv_d)
}

// OK Tested
// @Summary Create a new inventory details record
// @Description Add a new record to the inventory details table
// @Tags inventory_details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param inventory body models.InventoryDetails true "Inventory Details Data"
// @Success 201 {object} models.InventoryDetails
// @Failure 400 {object} map[string]string
// @Router /inventory-details [post]
func CreateInventoryDetails(w http.ResponseWriter, r *http.Request) {
	var invd models.InventoryDetails
	err := json.NewDecoder(r.Body).Decode(&invd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO inventory_details (transaction_id, product_id, transaction_type, quantity, transaction_description, transaction_amount, transaction_status, transaction_notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", invd.TransactionId, invd.ProductId, invd.TransactionType, invd.Quantity, invd.TransactionDescription, invd.TransactionAmount, invd.TransactionStatus, invd.TransactionNotes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Successfully Inserterd !")
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invd.TransactionId = int(id)
	now := "now"
	invd.TransactionDate = &now
	json.NewEncoder(w).Encode(invd)
}

// OK Tested
// @Summary Update inventory details record by ID
// @Description Modify an existing inventory details record
// @Tags inventory_details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param inventory body models.InventoryDetails true "Updated Inventory Details Data"
// @Success 200 {object} models.InventoryDetails
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /inventory-details/{id} [put]
func UpdateInventoryDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var invd models.InventoryDetails
	err := json.NewDecoder(r.Body).Decode(&invd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("UPDATE inventory_details SET product_id = ?, transaction_type = ?, quantity = ?, transaction_description = ?, transaction_amount = ?, transaction_status = ?, transaction_notes = ? WHERE transaction_id = ?", invd.ProductId, invd.TransactionType, invd.Quantity, invd.TransactionDescription, invd.TransactionAmount, invd.TransactionStatus, invd.TransactionNotes, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Successfully inventory_details Updated ID:", id)
	}
	json.NewEncoder(w).Encode(invd)
}

// OK Tested
// @Summary Delete inventory details record by ID
// @Description Remove an inventory details record from the database
// @Tags inventory_details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /inventory-details/{id} [delete]
func DeleteInventoryDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM inventory_details WHERE transaction_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
