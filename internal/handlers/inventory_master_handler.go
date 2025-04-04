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

/// Define API methods - Inventory Master

// OK Tested

// @Summary Get all inventory master records
// @Description Fetch all inventory master records from the database
// @Tags inventory_master
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Success 200 {array} models.InventoryMaster
// @Failure 500 {object} map[string]string
// @Router /inventory-master [get]
func GetAllInventoryMaster(w http.ResponseWriter, r *http.Request) {
	var inv_m []models.InventoryMaster
	rows, err := db.DB.Query("SELECT * FROM inventory_master")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Retrieve all inventry master data")
	}

	defer rows.Close()

	for rows.Next() {
		var invm models.InventoryMaster

		if err := rows.Scan(
			&invm.ProductId,
			&invm.ProductName,
			&invm.Category,
			&invm.Brand,
			&invm.StockInHand,
			&invm.UnitPrice,
			&invm.SupplierId,
			&invm.DateAdded,
			&invm.LastUpdated,
			&invm.Remarks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_m = append(inv_m, invm)
	}
	json.NewEncoder(w).Encode(inv_m)
}

// OK Tested
// @Summary Fetch inventory master record by ID
// @Description Get a single inventory master record using its ID
// @Tags inventory_master
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.InventoryMaster
// @Failure 404 {object} map[string]string
// @Router /inventory-master/{id} [get]
func GetInventoryMasterById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var invm models.InventoryMaster
	err := db.DB.QueryRow("SELECT * FROM inventory_master WHERE product_id = ?", id).Scan(
		&invm.ProductId,
		&invm.ProductName,
		&invm.Category,
		&invm.Brand,
		&invm.StockInHand,
		&invm.UnitPrice,
		&invm.SupplierId,
		&invm.DateAdded,
		&invm.LastUpdated,
		&invm.Remarks)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		fmt.Println("Get Specific Inventory Master Data")
	}
	json.NewEncoder(w).Encode(invm)
}

// OK Tested
// @Summary Filter inventory master records
// @Description Retrieve inventory master records based on query parameters
// @Tags inventory_master
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
// @Success 200 {array} models.InventoryMaster
// @Failure 500 {object} map[string]string
// @Router /inventory-master-filter [get]
func GetFilteredInventoryMaster(w http.ResponseWriter, r *http.Request) {

	query := "SELECT * FROM inventory_master WHERE 1=1"
	args := []interface{}{}

	if product_id := r.URL.Query().Get("product_id"); product_id != "" {
		query += " AND product_id = ?"
		args = append(args, product_id)
	}

	if product_name := r.URL.Query().Get("product_name"); product_name != "" {
		query += " AND product_name = ?"
		args = append(args, product_name)
	}

	if category := r.URL.Query().Get("category"); category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	if brand := r.URL.Query().Get("brand"); brand != "" {
		query += " AND brand = ?"
		args = append(args, brand)
	}

	if stock_in_hand := r.URL.Query().Get("stock_in_hand"); stock_in_hand != "" {
		query += " AND stock_in_hand = ?"
		args = append(args, stock_in_hand)
	}

	if unit_price := r.URL.Query().Get("unit_price"); unit_price != "" {
		query += " AND unit_price = ?"
		args = append(args, unit_price)
	}

	if supplier_id := r.URL.Query().Get("supplier_id"); supplier_id != "" {
		query += " AND supplier_id = ?"
		args = append(args, supplier_id)
	}

	if date_added := r.URL.Query().Get("date_added"); date_added != "" {
		query += " AND date_added = ?"
		args = append(args, date_added)
	}

	if last_updated := r.URL.Query().Get("last_updated"); last_updated != "" {
		query += " AND last_updated = ?"
		args = append(args, last_updated)
	}

	if remarks := r.URL.Query().Get("remarks"); remarks != "" {
		query += " AND remarks >= ?"
		args = append(args, remarks)
	}

	rows, err := db.DB.Query(query, args...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inv_m []models.InventoryMaster
	for rows.Next() {
		var invm models.InventoryMaster
		err := rows.Scan(&invm.ProductId, &invm.ProductName, &invm.Category, &invm.Brand,
			&invm.StockInHand, &invm.UnitPrice, &invm.SupplierId,
			&invm.DateAdded, &invm.LastUpdated, &invm.Remarks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inv_m = append(inv_m, invm)
	}

	json.NewEncoder(w).Encode(inv_m)
}

// OK Tested

// @Summary Create a new inventory master record
// @Description Add a new record to the inventory master table
// @Tags inventory_master
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param inventory body models.InventoryMaster true "Inventory Master Data"
// @Success 201 {object} models.InventoryMaster
// @Failure 400 {object} map[string]string
// @Router /inventory-master [post]
func CreateInventoryMaster(w http.ResponseWriter, r *http.Request) {
	var invm models.InventoryMaster
	err := json.NewDecoder(r.Body).Decode(&invm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO inventory_master (product_name, category, brand, stock_in_hand, unit_price, supplier_id, remarks) VALUES (?, ?, ?, ?, ?, ?, ?)", invm.ProductName, invm.Category, invm.Brand, invm.StockInHand, invm.UnitPrice, invm.SupplierId, invm.Remarks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Successfully Inserterd Inventory Master Data!")
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invm.ProductId = int(id)
	now := "now"
	invm.DateAdded = &now
	json.NewEncoder(w).Encode(invm)
}

// OK Tested

// @Summary Update inventory master record by ID
// @Description Modify an existing inventory master record
// @Tags inventory_master
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param inventory body models.InventoryMaster true "Updated Inventory Master Data"
// @Success 200 {object} models.InventoryMaster
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /inventory-master/{id} [put]
func UpdateInventoryMaster(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var invm models.InventoryMaster
	err := json.NewDecoder(r.Body).Decode(&invm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := "now"
	invm.LastUpdated = &now

	_, err = db.DB.Exec("UPDATE inventory_master SET product_name = ?, category = ?, brand = ?, stock_in_hand = ?, unit_price = ?, supplier_id = ?, last_updated = CURRENT_TIMESTAMP, remarks = ? WHERE product_id = ?", invm.ProductName, invm.Category, invm.Brand, invm.StockInHand, invm.UnitPrice, invm.SupplierId, invm.Remarks, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Successfully Updated Inventory Master ID:", id)
	}
	json.NewEncoder(w).Encode(invm)
}

// OK Tested
// @Summary Delete inventory master record by ID
// @Description Remove an inventory master record from the database
// @Tags inventory_master
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /inventory-master/{id} [delete]
func DeleteInventoryMaster(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM inventory_master WHERE product_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Successfully Deleted Inventory Master ID:", id)
	}
	w.WriteHeader(http.StatusNoContent)
}
