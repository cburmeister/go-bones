package main

import "net/http"
import "github.com/fengsp/knight"
import "github.com/gorilla/mux"
import "github.com/cburmeister/go-bones/routes"

func setupRoutes() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{username}", routes.User)
    http.Handle("/", r)
}

func setupReloader() {
    knight := knight.NewKnight("")
    knight.ListenAndServe(":8000", nil)
}

func main() {
    setupRoutes()
    setupReloader()
}
