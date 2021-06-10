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
func CreatePost(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // create an empty user of type models.User
    var post models.Posts

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&post)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertPost(post)

    // format a response object
    res := response{
        ID:      insertID,
        Message: "Post created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its ids
func GetPost(w http.ResponseWriter, r *http.Request) {
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
    post, err := getPost(int64(id))

    if err != nil {
        log.Fatalf("Unable to get user. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(post)
}

// GetAllUser will return all the users
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the users in the db
    posts, err := getAllPosts()

    if err != nil {
        log.Fatalf("Unable to get all user. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(posts)
}



//------------------------- handler functions ----------------
// insert one user in the DB
func insertPost(post models.Posts) string {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO posts (userid, message_txt) VALUES ($1, $2) RETURNING userid`

    // the inserted id will store in this id
    var id string

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, post.UserId, post.Message).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    // return the inserted id
    return id
}

// get one post from the DB by its userid
func getPost(id int64) (models.Posts, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var post models.Posts

    // create the select sql query
    sqlStatement := `SELECT * FROM posts WHERE postid=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to user
    err := row.Scan(&post.ID, &post.UserId, &post.Message)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return post, nil
    case nil:
        return post, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return post, err
}

func getAllPosts() ([]models.Posts, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var posts []models.Posts

    // create the select sql query
    sqlStatement := `SELECT * FROM posts`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()


    // iterate over the rows
    for rows.Next() {
        var post models.Posts

        // unmarshal the row object to user
        err = rows.Scan(&post.ID, &post.UserId, &post.Message)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        posts = append(posts, post)

    }

    // return empty user on error
    return posts, err
}	