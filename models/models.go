package models

// User schema of the user table
type User struct {
	ID         string `json:"userid"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Password   string `json:"password_hash"`
	IsAdmin    bool   `json:"isAdmin"`
}

// Posts schema of the posts table
type Posts struct {
	ID         int64 `json:"postid"`
	UserId     string `json:"userid"`
	Message    string `json:"message_txt"`
}

// Comments schema of the comments table
type Comments struct {
	ID         int64 `json:"commentid"`
	UserId     string `json:"userid"`
	PostId     int64 `json:"postid"`
	Message    string `json:"message_txt"`
}
