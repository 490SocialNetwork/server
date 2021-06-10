CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    isAdmin BOOLEAN DEFAULT FALSE
);
