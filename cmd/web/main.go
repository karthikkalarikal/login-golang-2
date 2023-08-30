package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthikkalarikal/logingolang2/pkg/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.LoginPage)
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/home", handlers.HomePage)
	r.HandleFunc("/signout", handlers.LogoutHandler)

	fmt.Println("Server is listening on :8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
