package routes

import "net/http"
import "html/template"
import "github.com/gorilla/mux"
import "github.com/cburmeister/go-bones/lib"
import "github.com/cburmeister/go-bones/models"

func User(w http.ResponseWriter, r *http.Request) {
    lib.Log(r.RemoteAddr, r.Method, r.URL.Path)

    vars := mux.Vars(r)

    user := &models.User{
	Username: vars["username"],
    }

    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(w, user)
}
