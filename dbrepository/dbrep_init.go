package dbrepository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DBInit - создает соединение с базой данных
func DBInit() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbName := os.Getenv("db_name")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	log.Println(dbUri)
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
