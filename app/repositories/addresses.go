package main

import (
		"../schemas"
		"database/sql"
		"fmt"
	)

func GetAddresses(db *sql.DB, id int) schemas.Adress {
	query := fmt.Sprintf(`SELECT * FROM address WHERE id = %d;`, id)

	var adress schemas.Address

	row := db.QueryRow(query)

	switch err := row.Scan(&address.Adress1, &address.Adress2, &address.City, &address.State,
		&address.Country, &address.ZipCode, &address.MerchantId); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(address)
	default:
		panic(err)
	}

	return adress
}

func GetAllAddresses(db *sql.DB) []schemas.adress {
	query := fmt.Sprintf(`SELECT * FROM address`)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}
	var address []schemas.address

	for rows.Next() {
		var adress schemas.address
		err = rows.Scan(
			&address.Adress1, &address.Adress2, &address.City, &address.State,
			&address.Country, &address.ZipCode, &address.MerchantId)
		address = append(address, address)
	}

	if err != nil {
		panic(err)
	}

	return address
}

func InsertAdress(db *sql.DB, adress schemas.adress) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT INTO address(address1, address2, city, state, country, ZipCode, MerchantId) VALUES ('%s','%s','%s','%s','%s','%s','%s')`,
		address.Adress1, address.Adress2, address.City, address.State,
		address.Country, address.ZipCode, address.MerchantId)

	result, err := db.Exec(query)

	return result, err
}
