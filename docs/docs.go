// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/inventory-details": {
            "get": {
                "description": "Get a list of all inventory details records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Fetch all inventory details records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryDetails"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new record to the inventory details table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Create a new inventory details record",
                "parameters": [
                    {
                        "description": "Inventory Details Data",
                        "name": "inventory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InventoryDetails"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryDetails"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-details-filter": {
            "get": {
                "description": "Retrieve inventory details records based on query parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Filter inventory details records",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Type",
                        "name": "transaction_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Quantity",
                        "name": "quantity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Date",
                        "name": "transaction_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Description",
                        "name": "transaction_description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Amount",
                        "name": "transaction_amount",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Status",
                        "name": "transaction_status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Notes",
                        "name": "transaction_notes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryDetails"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-details-overall": {
            "get": {
                "description": "Get a comprehensive list of inventory and transaction details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_overall"
                ],
                "summary": "Fetch overall inventory details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryOverall"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-details-overall-filter": {
            "get": {
                "description": "Retrieve overall inventory records based on query parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_overall"
                ],
                "summary": "Filter overall inventory details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Product Name",
                        "name": "product_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Brand",
                        "name": "brand",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Stock In Hand",
                        "name": "stock_in_hand",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Unit Price",
                        "name": "unit_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Supplier ID",
                        "name": "supplier_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Date Added",
                        "name": "date_added",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Last Updated",
                        "name": "last_updated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Remarks",
                        "name": "remarks",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Type",
                        "name": "transaction_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Quantity",
                        "name": "quantity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Date",
                        "name": "transaction_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Description",
                        "name": "transaction_description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Amount",
                        "name": "transaction_amount",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Status",
                        "name": "transaction_status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction Notes",
                        "name": "transaction_notes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryOverall"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-details/{id}": {
            "get": {
                "description": "Get a single inventory details record using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Fetch inventory details by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Modify an existing inventory details record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Update inventory details record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Inventory Details Data",
                        "name": "inventory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InventoryDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryDetails"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an inventory details record from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_details"
                ],
                "summary": "Delete inventory details record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-master": {
            "get": {
                "description": "Fetch all inventory master records from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Get all inventory master records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryMaster"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new record to the inventory master table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Create a new inventory master record",
                "parameters": [
                    {
                        "description": "Inventory Master Data",
                        "name": "inventory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InventoryMaster"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryMaster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-master-filter": {
            "get": {
                "description": "Retrieve inventory master records based on query parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Filter inventory master records",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Product Name",
                        "name": "product_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Brand",
                        "name": "brand",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Stock In Hand",
                        "name": "stock_in_hand",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Unit Price",
                        "name": "unit_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Supplier ID",
                        "name": "supplier_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Date Added",
                        "name": "date_added",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Last Updated",
                        "name": "last_updated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Remarks",
                        "name": "remarks",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.InventoryMaster"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/inventory-master/{id}": {
            "get": {
                "description": "Get a single inventory master record using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Fetch inventory master record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryMaster"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Modify an existing inventory master record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Update inventory master record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Inventory Master Data",
                        "name": "inventory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InventoryMaster"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InventoryMaster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an inventory master record from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory_master"
                ],
                "summary": "Delete inventory master record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Credentials Data",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Credentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Credentials"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Credentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.InventoryDetails": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "transaction_amount": {
                    "type": "number"
                },
                "transaction_date": {
                    "type": "string"
                },
                "transaction_description": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "integer"
                },
                "transaction_notes": {
                    "type": "string"
                },
                "transaction_status": {
                    "type": "string"
                },
                "transaction_type": {
                    "type": "string"
                }
            }
        },
        "models.InventoryMaster": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "date_added": {
                    "type": "string"
                },
                "last_updated": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "stock_in_hand": {
                    "type": "integer"
                },
                "supplier_id": {
                    "type": "integer"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        },
        "models.InventoryOverall": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "date_added": {
                    "type": "string"
                },
                "last_updated": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "remarks": {
                    "type": "string"
                },
                "stock_in_hand": {
                    "type": "integer"
                },
                "supplier_id": {
                    "type": "integer"
                },
                "transaction_amount": {
                    "type": "number"
                },
                "transaction_date": {
                    "type": "string"
                },
                "transaction_description": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "integer"
                },
                "transaction_notes": {
                    "type": "string"
                },
                "transaction_status": {
                    "type": "string"
                },
                "transaction_type": {
                    "type": "string"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Inventory Management API",
	Description:      "This is an API for managing inventory",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
