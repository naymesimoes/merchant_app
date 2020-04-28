CREATE TABLE address(
    id SERIAL PRIMARY KEY,
    address1 varchar,
    address2 varchar,
    city varchar,
    state varchar,
    country varchar,
    zip_code varchar,
    merchant_id int,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
)
