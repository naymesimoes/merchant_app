package main

import(
	"../schemas"
	"database/sql"
	"fmt"
)

func GetAllMerchants(db *sql.DB) (schemas.Merchant){
	query := `SELECT id, name, last_name FROM merchants;`

	var merchant schemas.Merchant;

	row := db.QueryRow(query)

	switch err := row.Scan(&merchant.Id, &merchant.Name, &merchant.LastName); err {
		case sql.ErrNoRows:
  			fmt.Println("No rows were returned!")
		case nil:
  			fmt.Println(merchant)
		default:
  			panic(err)
	}

	return merchant;
}
