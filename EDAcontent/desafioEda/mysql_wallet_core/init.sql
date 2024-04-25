-- Inits walletCore database container with this migrations. 
CREATE TABLE clients (
    id VARCHAR(255), 
    name VARCHAR(255),
    email VARCHAR(255), 
    created_at date,
    PRIMARY KEY(id)    
);

CREATE TABLE accounts (
    id VARCHAR(255), 
    client_id VARCHAR(255), 
    balance int, 
    created_at date,
    PRIMARY KEY(id)
);

CREATE TABLE transactions (
    id VARCHAR(255),
    account_id_from VARCHAR(255),
    account_id_to VARCHAR(255), 
    amount int, 
    created_at date,
    PRIMARY KEY(id)
); 


-- Jhon
INSERT INTO clients (id, name, email, created_at) VALUES (
    "8f604540-cf40-44c6-9584-16e5757c9695",
    "jhon doe", "j@j.com", null
);

-- Jane 
INSERT INTO clients (id, name, email, created_at) VALUES (
    "7fa0698f-143d-4edd-b00a-2b12dff9b6ff",
    "Jane doe", "jane@j.com", null
);

-- Jhon Account
INSERT INTO accounts (id, client_id, balance, created_at) VALUES (
    "02c08bd5-7ec2-45d9-8f27-8c3422927b6a", 
    "8f604540-cf40-44c6-9584-16e5757c9695", 
    1000,
    null
);

-- Jane Account 
INSERT INTO accounts (id, client_id, balance, created_at) VALUES (
    "c98ed461-eb0c-47e3-b142-b5fbdf86dc41", 
    "7fa0698f-143d-4edd-b00a-2b12dff9b6ff", 
    1000,
    null
);

-- Use software, to create transactions (need to send msg in kafka)