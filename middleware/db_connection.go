package middleware

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv" // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
    ID      string  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

type response_int struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}


// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}