CREATE TABLE posts (
    postid SERIAL PRIMARY KEY,
    userid TEXT REFERENCES users (userid),
    message_txt TEXT NOT NULL
);