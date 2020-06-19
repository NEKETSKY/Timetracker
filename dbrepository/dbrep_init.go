package dbrepository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DBInit - allows you to create a database connection
func DBInit() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbUri := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbHost, dbName, dbPort)
	log.Println(dbUri)
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
