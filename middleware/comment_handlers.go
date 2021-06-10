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
func CreateComment(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // create an empty user of type models.User
    var comment models.Comments

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&comment)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertComment(comment)

    // format a response object
    res := response{
        ID:      insertID,
        Message: "User created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetComment(w http.ResponseWriter, r *http.Request) {
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
    comment, err := getComment(int64(id))

    if err != nil {
        log.Fatalf("Unable to get user. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(comment)
}

// GetAllUser will return all the users
func GetAllComments(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the users in the db
    comments, err := getAllComments()

    if err != nil {
        log.Fatalf("Unable to get all user. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(comments)
}



//------------------------- handler functions ----------------
// insert one user in the DB
func insertComment(comment models.Comments) int64 {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO comments (userid,postid, message_txt) VALUES ($1, $2, $3) RETURNING userid`

    // the inserted id will store in this id
    var id int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, comment.UserId, comment.PostId, comment.Message).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    // return the inserted id
    return id
}

// get one post from the DB by its userid
func getComment(id int64) (models.Comments, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var comment models.Comments

    // create the select sql query
    sqlStatement := `SELECT * FROM comments WHERE commentid=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to user
    err := row.Scan(&comment.ID, &comment.UserId, &comment.PostId, &comment.Message)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return comment, nil
    case nil:
        return comment, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return comment, err
}

func getAllComments() ([]models.Comments, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var comments []models.Comments

    // create the select sql query
    sqlStatement := `SELECT * FROM comments`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()


    // iterate over the rows
    for rows.Next() {
        var comment models.Comments

        // unmarshal the row object to user
        err = rows.Scan(&comment.ID, &comment.UserId, &comment.PostId, &comment.Message)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        comments = append(comments, comment)

    }

    // return empty user on error
    return comments, err
}	