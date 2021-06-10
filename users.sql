CREATE TABLE users (
    userid TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    isAdmin BOOLEAN DEFAULT FALSE
);
