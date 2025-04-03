package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:root@tcp(localhost:3306)/inventory"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	
	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}
	fmt.Println("Connected to the database successfully")
}
