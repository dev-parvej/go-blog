package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	"github.com/dev-parvej/go-blog/middleware"
	route "github.com/dev-parvej/go-blog/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Blog with go")
	r.HandleFunc("/", route.ServeHome)
	r.HandleFunc("/sign-up", route.CreateUserForm).Methods("GET")
	r.HandleFunc("/register", route.CreateUser).Methods("POST")
	r.HandleFunc("/sign-in", route.LoginForm).Methods("GET")
	r.HandleFunc("/login", route.Login).Methods("POST")
	r.HandleFunc("/sign-out", func(w http.ResponseWriter, r *http.Request) {
		var session = config.Session(r)
		session.Values["authenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusAccepted)
	}).Methods("GET")

	//Authenticated routes
	ar := r.PathPrefix("/posts").Subrouter()
	ar.Use(middleware.IsAuthenticated)
	ar.HandleFunc("/save", route.CreatePost).Methods("POST")
	ar.HandleFunc("/new", route.CreatePostForm).Methods("GET")
	ar.HandleFunc("/{id:[0-9]+}", route.Post).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Get("APP_PORT")), r))
}
