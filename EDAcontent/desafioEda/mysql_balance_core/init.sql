-- Inits balancecore database container with this migrations. 

CREATE TABLE balances (
    id VARCHAR(255), 
    account_id_from VARCHAR(255),
    account_id_to VARCHAR(255),
    balance_account_id_from int,
    balance_account_id_to int,
    created_at date,
    PRIMARY KEY(id)  
);

CREATE TABLE accounts (
    id VARCHAR(255),
    balance int, 
    PRIMARY KEY(id)
); 

-- Jhon Account
INSERT INTO accounts (id, balance) VALUES (
    "02c08bd5-7ec2-45d9-8f27-8c3422927b6a", 1000
);

-- Jane Account 
INSERT INTO accounts (id, balance) VALUES (
    "c98ed461-eb0c-47e3-b142-b5fbdf86dc41", 1000
);