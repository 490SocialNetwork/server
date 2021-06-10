package router

import (
    "go-postgres/middleware"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //Users
    router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
    //router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
    //router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")


    //Posts
    router.HandleFunc("/api/posts/{id}", middleware.GetPost).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/posts", middleware.GetAllPosts).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newpost", middleware.CreatePost).Methods("POST", "OPTIONS")

    //Comments
    router.HandleFunc("/api/comments/{id}", middleware.GetComment).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/comments", middleware.GetAllComments).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newcomment", middleware.CreateComment).Methods("POST", "OPTIONS")


    return router
}
