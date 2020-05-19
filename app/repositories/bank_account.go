package repositories

import (
	"database/sql"
	"fmt"

	"../schemas"
)

func GetBankAccountById(id int) schemas.BankAccount {
	query := fmt.Sprintf(`SELECT * FROM bank_account WHERE id = %d;`, id)

	var bank_account schemas.BankAccount

	row := db.QueryRow(query)

	switch err := row.Scan(&bank_account.Id, &bank_account.BranchCode,
		&bank_account.AccountNumber, &bank_account.Type,
		&bank_account.BankCode, &bank_account.MerchantId); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(bank_account)
	default:
		panic(err)
	}

	return bank_account
}

func GetAllBankAccountes() []schemas.BankAccount {
	query := fmt.Sprintf(`SELECT * FROM bank_account`)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}
	var bank_accountes []schemas.BankAccount

	for rows.Next() {
		var bank_account schemas.BankAccount
		err = rows.Scan(&bank_account.Id, &bank_account.BranchCode,
			&bank_account.AccountNumber, &bank_account.Type,
			&bank_account.BankCode, &bank_account.MerchantId)
		bank_accountes = append(bank_accountes, bank_account)
	}

	if err != nil {
		panic(err)
	}

	return bank_accountes
}

func InsertBankAccount(bank_account schemas.BankAccount) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT INTO bank_account(branch_code, account_number, type, bank_code, merchant_id) VALUES ('%s','%s','%s',%d,%d)`,
		bank_account.BranchCode,
		bank_account.AccountNumber, bank_account.Type,
		bank_account.BankCode, bank_account.MerchantId)

	result, err := db.Exec(query)

	return result, err
}
