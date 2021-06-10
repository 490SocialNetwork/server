CREATE TABLE comments (
    commentid SERIAL PRIMARY KEY,
    userid TEXT REFERENCES users (userid),
    postid INTEGER REFERENCES posts (postid),
    message_txt TEXT NOT NULL
);