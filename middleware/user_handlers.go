package middleware

import (
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "go-postgres/models" // models package where User schema is defined
    "log"
    "net/http" // used to access the request and response object of the api
    "strconv"  // package used to covert string into int type

    "github.com/gorilla/mux" // used to get the params from the route
    _ "github.com/lib/pq"      // postgres golang driver
)

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // create an empty user of type models.User
    var user models.User

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertUser(user)

    // format a response object
    res := response{
        ID:      insertID,
        Message: "User created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // call the getUser function with user id to retrieve a single user
    user, err := getUser(int64(id))

    if err != nil {
        log.Fatalf("Unable to get user. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the users in the db
    users, err := getAllUsers()

    if err != nil {
        log.Fatalf("Unable to get all user. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(users)
}

//------------------------- handler functions ----------------
// insert one user in the DB
func insertUser(user models.User) string {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO users (userid,first_name, last_name, password_hash) VALUES ($1, $2, $3, $4) RETURNING userid`

    // the inserted id will store in this id
    var id string

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, user.ID, user.First_Name, user.Last_Name, user.Password).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    // return the inserted id
    return id
}

// get one user from the DB by its userid
func getUser(id int64) (models.User, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var user models.User

    // create the select sql query
    sqlStatement := `SELECT * FROM users WHERE userid=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to user
    err := row.Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Password, &user.IsAdmin)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return user, nil
    case nil:
        return user, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return user, err
}

func getUserLoginInfo(email string) (models.User, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var user models.User

    // create the select sql query
    sqlStatement := `SELECT email,password_hash FROM users WHERE email=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, email)

    // unmarshal the row object to user
    err := row.Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Password, &user.IsAdmin)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return user, nil
    case nil:
        return user, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return user, err
}

// get one user from the DB by its userid
func getAllUsers() ([]models.User, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var users []models.User

    // create the select sql query
    sqlStatement := `SELECT * FROM users`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()


    // iterate over the rows
    for rows.Next() {
        var user models.User

        // unmarshal the row object to user
        err = rows.Scan(&user.ID, &user.First_Name, &user.Last_Name, &user.Password, &user.IsAdmin)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        users = append(users, user)

    }

    // return empty user on error
    return users, err
}	