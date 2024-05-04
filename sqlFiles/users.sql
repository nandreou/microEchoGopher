CREATE TABLE users (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    username VARCHAR(255) UNIQUE,
    pass VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);


INSERT INTO users(first_name, last_name, email, username, pass) 
VALUES ('nick', 'test', 'email@email.com', 'nick1', 'pass1'),
       ('john', 'doe', 'john.doe@example.com', 'nick2', 'pass2'),
       ('jane', 'smith', 'jane.smith@example.com', 'nick3', 'pass3');