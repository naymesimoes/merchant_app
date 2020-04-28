CREATE TABLE bank_account(
    id SERIAL PRIMARY KEY,
    branch_code varchar,
    account_number varchar,
    type varchar,
    bank_code int,
    merchant_id int,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
)
