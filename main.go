package main

import "net/http"
import "fmt"
import "github.com/fengsp/knight"
import "github.com/gorilla/mux"

func log(remote_addr string, method string, path string) {
    fmt.Println(remote_addr, method, path)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]

    fmt.Fprintf(w, "Hi there %s!", username)
}

func main() {
    setup_routes()
    setup_reloader()
}

func setup_routes() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{username}", UserHandler)
    http.Handle("/", r)
}

func setup_reloader() {
    knight := knight.NewKnight("")
    knight.ListenAndServe(":8000", nil)
}
