package main

import (
    "go-postgres/router"
    "log"
    "net/http"
    "os"
)

func main() {
    r := router.Router()
    var port string
    var def bool
    if port, def = os.LookupEnv("PORT"); !def {
        port = "8000"
    }
    log.Fatal(http.ListenAndServe(":",port, r))
}
