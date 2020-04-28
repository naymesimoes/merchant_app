package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/golang-migrate/migrate/v4"
    // "github.com/golang-migrate/migrate/v4/database/postgres"
    // _ "github.com/golang-migrate/migrate/v4/source/file"
	"fmt"
	"../schemas"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "merchant_app"
  )

var db *sql.DB

func main() {
	
	db, _ = Conection()

	// GetMerchantById(db, 1);

	// values := GetAllMerchants(db);

	// for i := 0; i<2; i++ {
	// 	fmt.Println(values[i])
	// }
	// value := schemas.Merchant{3,"Nayme","Simoes","","13213213290","1234-1234","98765-1234","nay_s@email.com"}
	// InsertMerchant(db, value)

	value := schemas.BankAccount{0,"0123", "130038465", "cc", 100,3}

	result,err := InsertBankAccount(value)

	fmt.Println(result)
	fmt.Println(err)

	CloseConection()
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

func CloseConection() {
	db.Close()
}
