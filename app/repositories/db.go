package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
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

	// migration test
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
        "../../migrations",
		"postgres", driver)
	m.Steps(2)
	// end migration test

	// SELECT
	query := `SELECT id, serial_number FROM card_readers;`

	var id int
	var serial_number string

	row := db.QueryRow(query)

	switch err := row.Scan(&id, &serial_number); err {
		case sql.ErrNoRows:
  			fmt.Println("No rows were returned!")
		case nil:
  			fmt.Println(id, serial_number)
		default:
  			panic(err)
	}
	// END SELECT
	
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

