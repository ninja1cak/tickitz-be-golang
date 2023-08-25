package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PgDatabase() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbName)

	return sqlx.Connect("postgres", config)
}
