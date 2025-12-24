CREATE TABLE roles (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);