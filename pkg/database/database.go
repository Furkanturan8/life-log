package database

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL sürücüsü
)

var db *bun.DB

func StartPostgresDB() error {
	// Bağlantı dizesini oluştur
	connStr := os.Getenv("POSTGRES_URI")
	if connStr == "" {
		return fmt.Errorf("missing 'POSTGRES_URI' environment variable")
	}

	// Veritabanına bağlan
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db = bun.NewDB(sqlDB, pgdialect.New())

	// Bağlantıyı test et
	if err = db.Ping(); err != nil {
		return fmt.Errorf("database connection test error: %w", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	return nil
}

func ClosePostgresDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error closing database connection: %v\n", err)
		} else {
			fmt.Println("Database connection closed.")
		}
	}
}

func GetDB() *bun.DB {
	return db
}
