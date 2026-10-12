-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    price DOUBLE PRECISION NOT NULL,
    image TEXT
);