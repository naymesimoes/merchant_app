package repositories

import (
	"database/sql"
	"fmt"

	"../schemas"
)

func GetMerchantById(id int) schemas.Merchant {
	query := fmt.Sprintf(`SELECT * FROM merchants WHERE id = %d;`, id)

	var merchant schemas.Merchant

	row := db.QueryRow(query)

	switch err := row.Scan(&merchant.Id, &merchant.Name,
		&merchant.LastName, &merchant.SignupedAt,
		&merchant.Cpf, &merchant.PhoneNumber1,
		&merchant.PhoneNumber2, &merchant.Email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(merchant)
	default:
		panic(err)
	}

	return merchant
}

func GetAllMerchants() []schemas.Merchant {
	query := fmt.Sprintf(`SELECT * FROM merchants`)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}
	var merchants []schemas.Merchant

	for rows.Next() {
		var merchant schemas.Merchant
		err = rows.Scan(
			&merchant.Id, &merchant.Name,
			&merchant.LastName, &merchant.SignupedAt,
			&merchant.Cpf, &merchant.PhoneNumber1,
			&merchant.PhoneNumber2, &merchant.Email)
		merchants = append(merchants, merchant)
	}

	if err != nil {
		panic(err)
	}

	return merchants
}

func InsertMerchant(merchant schemas.Merchant) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT INTO merchants(id,name,last_name,signuped_at,cpf,
		phone_number1,phone_number2,email) VALUES (%d,'%s','%s',now(),'%s','%s','%s','%s')`,
		merchant.Id, merchant.Name, merchant.LastName, merchant.Cpf,
		merchant.PhoneNumber1, merchant.PhoneNumber2, merchant.Email)

	result, err := db.Exec(query)

	return result, err
}
