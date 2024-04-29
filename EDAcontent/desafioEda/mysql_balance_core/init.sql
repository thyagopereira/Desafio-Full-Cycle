-- Inits balancecore database container with this migrations. 

CREATE TABLE balances (
    id VARCHAR(255), 
    account_id VARCHAR(255),
    amount int, 
    created_at date,
    PRIMARY KEY(id)  
);

