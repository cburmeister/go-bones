package main

import "net/http"
import "fmt"
import "github.com/fengsp/knight"
import "github.com/gorilla/mux"

func log(remote_addr string, method string, path string) {
    fmt.Println(remote_addr, method, path)
}

type loggedResponse struct {
    w http.ResponseWriter
    r *http.Request
    url string
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]

    fmt.Fprintf(w, "Hi there %s!", username)

    logged_response := &loggedResponse{w, r, r.URL.String()}

    log(r.RemoteAddr, r.Method, r.URL.Path)
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
