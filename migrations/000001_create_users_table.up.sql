CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(75) NOT NULL,
    password_hash TEXT NOT NULL
)