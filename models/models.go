package models

// User schema of the user table
type User struct {
	ID         int64  `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password_hash"`
	IsAdmin    bool   `json:"isAdmin"`
}
