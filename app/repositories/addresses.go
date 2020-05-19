package repositories

import (
	"database/sql"
	"fmt"

	"../schemas"
)

func GetAddressById(id int) schemas.Address {
	query := fmt.Sprintf(`SELECT * FROM address WHERE id = %d;`, id)

	var address schemas.Address

	row := db.QueryRow(query)

	switch err := row.Scan(&address.Id, &address.Address1, &address.Address2, &address.City, &address.State,
		&address.Country, &address.ZipCode, &address.MerchantId); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(address)
	default:
		panic(err)
	}

	return address
}

func GetAllAddresses() []schemas.Address {
	query := fmt.Sprintf(`SELECT * FROM address`)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}
	var addresses []schemas.Address

	for rows.Next() {
		var address schemas.Address
		err = rows.Scan(
			&address.Id, &address.Address1, &address.Address2, &address.City, &address.State,
			&address.Country, &address.ZipCode, &address.MerchantId)
		addresses = append(addresses, address)
	}

	if err != nil {
		panic(err)
	}

	return addresses
}

func InsertAddress(address schemas.Address) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT INTO address(address1, address2, city, state, country, zip_code, merchant_id) VALUES ('%s','%s','%s','%s','%s','%s',%d)`,
		address.Address1, address.Address2, address.City, address.State,
		address.Country, address.ZipCode, address.MerchantId)

	result, err := db.Exec(query)

	return result, err
}
