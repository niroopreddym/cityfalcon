GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;

-- lookup tables and the static data insertion part
CREATE TABLE Bank (
    bank_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    bank_uuid uuid,
    bank_name varchar(50),
    ifsc_code varchar(50) unique,
    branch_name varchar(50)
);

CREATE TABLE Account (
    account_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    account_uuid uuid,
    account_holder_name varchar(50),
    bank_id int,
    first_name varchar(50),
    last_name varchar(50),
    balance varchar(50),
    Constraint fk_bank foreign key(bank_id) references bank(bank_id)
);