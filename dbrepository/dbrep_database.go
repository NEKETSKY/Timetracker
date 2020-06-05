package dbrepository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var dbase *sql.DB

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
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}

// возвращает дескриптор объекта DB
func GetDB() *sql.DB {
	return dbase
}

//функция - костяк под переделку, она не возвращает структуру тасков и вообще нужно переделать, потому что будет выводить в другую структуру
func GetGroups() []Group {

	xuilo, _:=GetDB().Query("SELECT * FROM groups;") //ОБРАБОТКУ ОШИБОК ДОБАВИТЬ
	defer xuilo.Close()
	var Groups []Group
	for xuilo.Next() {
		var group Group
		xuilo.Scan(&group.GroupID, &group.GroupTitle)
		Groups = append(Groups, group)
	}
	return Groups
}



/*
func GetGroups() []Group {

	xuilo, _ := dbase.Query("SELECT * FROM groups;") //ОБРАБОТКУ ОШИБОК ДОБАВИТЬ
	defer xuilo.Close()
	var Groups []Group
	for xuilo.Next() {
		var group Group
		xuilo.Scan(&group.GroupID, &group.GroupTitle)
		Groups = append(Groups, group)
	}
	return Groups
}
*/