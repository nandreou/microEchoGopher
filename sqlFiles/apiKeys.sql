CREATE TABLE apikeys(
    apikey VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255),
    created_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY (email) REFERENCES users(email)
);