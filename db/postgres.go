package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

func Init() *sql.DB {
	return getConnect()
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func getConnect() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("POSTGRES_USER", "postgres"),
		getEnv("POSTGRES_PASS", "postgres"),
		getEnv("POSTGRES_DBNAME", "sample"))
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	if os.Getenv("ENABLE_ELASTIC_APM") == "true" {
		fmt.Println("ELASTIC ENABLED")
		db, err = apmsql.Open("postgres", dbinfo)
	}
	return db
}
