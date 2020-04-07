package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/golang-migrate/migrate/v4"
    // "github.com/golang-migrate/migrate/v4/database/postgres"
    // _ "github.com/golang-migrate/migrate/v4/source/file"
	"fmt"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "merchant_app"
  )

func main() {
	
	db, _ := Conection()

	GetAllMerchants(db);
	
	CloseConection(db)
}
 
func Conection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, err
}

func CloseConection(db *sql.DB) {
	db.Close()
}
