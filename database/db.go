package database

import (
	"database/sql"
	"first_project/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitializeDB establishes a connection to PostgreSQL
func InitializeDB() {
	dbConfig := config.GetDBConfig()

	// Connection string (Data Source Name - DSN) format for PostgreSQL
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.SSLMode,
	)

	// Open the database connection
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening DB connection: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the PostgreSQL database successfully!")
}
