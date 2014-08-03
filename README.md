go-bones
===========

An example Golang web server.

---

Get a handful of dependencies:

```bash
$ go get "github.com/fengsp/knight"
$ go get "github.com/gorilla/mux"
```

Run the server:

```bash
$ go run main.go
* Knight serving on :8000
* Restarting with reloader
```

GET the `/user` endpoint:

```bash
$ curl localhost:8000/user/cburmeister
<h1>cburmeister</h1>
```
